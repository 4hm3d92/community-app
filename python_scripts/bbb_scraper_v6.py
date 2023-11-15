#!/usr/bin/python3
# A script for scraping bbb.org following their robots.txt rules
# It accepts as an argument the filename where links of businesses are stored
# Another script must be used to get the links of business from the sitemap of bbb.org

import asyncpg
import httpx
import json
import random
import asyncio
import sys
from datetime import datetime

profiles_file = sys.argv[1]
with open(profiles_file, 'r') as file:
    urls = [l for l in (line.strip() for line in file) if l]

print('Started on: ',datetime.now())

with open("useragent_list_8",'r') as file:
    useragentlist = file.read().splitlines()

useragents = []
for useragent in useragentlist:
    useragents.append({'user-agent': useragent})

with open(f"proxy_list", 'r') as file:
    proxylist = file.read().splitlines()

proxies = []
for proxy in proxylist:
    proxy = proxy.split(':')
    proxies.append({'proxy':{'all://': f'http://{proxy[2]}:{proxy[3]}@{proxy[0]}:{proxy[1]}'}, 'UA': random.choice(useragents)})

N=len(proxies)
#N=20

limits = httpx.Limits(max_connections=3,max_keepalive_connections=1)

timeout = httpx.Timeout(10, read=60)

async def parseProfile(n,url,pool,response):

    i = 3
    while i>=0:
        try:
            text = response.text
            if response.status_code == 403:
                return 2
            data = json.loads(text[text.find('{"header"'):text.rfind('}};</')+2])
            break
        except Exception as e:
            await asyncio.sleep(0.3)
            #print(n)
            print(e)
            print(n, 'HTTP', response.status_code, url)
            if response.status_code == 502:
                return 2
        
        i-=1
        if i==0:
            print('failed url: ',url)
            return 1

    profile = data["businessProfile"]
    contactInfo = profile["contactInformation"]

    businessId = profile["businessId"]
    businessName = profile["names"]["primary"]
    
    # DO NOT REPLACE GET. If get isn't used then a KeyError exception would occur wherever a profile that doesn't have the primary key is encountered.

    website = profile["urls"].get('primary')
    emails = {contactInfo["emailAddress"][8:-8].replace('__at__','@').replace('__dot__','.')} if contactInfo["emailAddress"] else set()
    
    for email in contactInfo["additionalEmailAddresses"]:
        if email:
            emails.add(email["value"][8:-8].replace('__at__','@').replace('__dot__','.'))
    
    #emails = ','.join(emails)
    emails = list(emails)

    phones = {contactInfo["phoneNumber"]} if contactInfo["phoneNumber"] else set()
    
    for phone in contactInfo["additionalPhoneNumbers"]:
        if phone:
            phones.add(phone["value"])
    
    phones = list(phones)
    #phones = ','.join(phones)
    
    postalAddress=profile["location"]["postalAddress"]
    zip = postalAddress["zipCode"]

    address = postalAddress['addressLine1']
    city = postalAddress.get('city')
    
    contacts = []
    for contact in contactInfo["contacts"]:
        name = contact['name']
        # contact will be title,prefix,firstname,lastname for eg. if lastname and title are missing, it will be stored as: ,Mr.,John,
        c = [contact['title'] if contact['title'] else '', name['prefix'] if name['prefix'] else '' , name['first'] if name['first'] else '' , name['last'] if name['last'] else '']
    
        if c:
            contacts.append(','.join(c))
    
    #contacts = ','.join(contacts)
    
    country = url[20:22].upper()
    state = url[23:25].upper()
    
    categories = []
    for category in profile['categories']['links']:
        categories.append(category['title'])
    
    await pool.execute('''
        INSERT INTO bbb_profile (business_id, business_name, website, profile_link, emails, phones, address, zip, state, country, categories, contacts, city) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) ON CONFLICT DO NOTHING;
                       ''', businessId, businessName, website, url, emails, phones, address, zip, state, country, categories, contacts, city)

    
    return 0


async def worker(n,queue,pool):
    async with httpx.AsyncClient(headers=proxies[n]['UA'],proxies=proxies[n]['proxy'],timeout=timeout,limits=limits,verify=False,http2=True,follow_redirects=True) as client:
        while True:
            try:
                url = queue.get_nowait()
            except asyncio.QueueEmpty:
                return

            try:
                
                response = await client.get(url)
                await asyncio.sleep(2)
                r = await parseProfile(n,url,pool,response)

            except Exception as e:
                print(n, e, url)
                r = 2
            

            # Delay between requests. 5 sec was probably getting proxies blocked.
            await asyncio.sleep(8)
            
            if r==0:
                print(f"Worker {n} scraped {url}")
            elif r==1:
                print('Url:', url, 'failed on worker:', n)
            elif r==2:
                print('Proxy:', proxies[n]['proxy'], 'might be blocked. Worker:', n, 'Url:', url)

                queue.put_nowait(url)

                # Since the url will be requeued, set this task as done, before breaking the loop and stopping the worker
                queue.task_done()
                break
                #clients[n]
                #return
            elif r!=0:
                while r!=0:
                    await asyncio.sleep(2)
            queue.task_done()


async def main():
    q=asyncio.Queue()
    print('Started on: ',datetime.now())
    
    for url in urls:
        q.put_nowait(url)
    
    async with asyncpg.create_pool(database="xxx", user='xxx',password='xxx', host='xxx', 
                               command_timeout=60, max_size=95) as pool:
        await asyncio.gather(*[worker(n,q,pool) for n in range(N)])
        await q.join()
    
    print('Finished on:', datetime.now())

if __name__ == "__main__":
    asyncio.run(main())

