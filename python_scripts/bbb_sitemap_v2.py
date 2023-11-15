import os
#import requests
import httpx
#import xml.etree.ElementTree as ET
from lxml import etree as ET

proxies = {
        'all://' : 'http://xxx:xxx@xxx:xxx'
}

headers = {
    'user-agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36'
}

#sitemap = 'https://www.bbb.org/sitemap-accredited-business-profiles-index.xml'
sitemap = 'https://www.bbb.org/sitemap-business-profiles-index.xml'

ns = {"bbb":"http://www.sitemaps.org/schemas/sitemap/0.9"}

profiles_dict = {}

def get_data(url):
    
    parser = ET.XMLParser(encoding="utf-8", recover=True)

    response = client.get(url)
    #print(response.text)
    root = ET.fromstring(bytes(response.text,encoding="utf-8"), parser=parser)
    links = []
    
    for link in root.findall('.//bbb:loc',ns):
        links.append(link.text)

    return links
    
#ca_profiles = []
#us_profiles = []
with httpx.Client(headers=headers,proxies=proxies,timeout=None) as client:
    #for sitemap in sitemaps:
    print("Processing "+sitemap)
    pages = get_data(sitemap)
    #pages = get_data(sitemap)[175:]
    for index in pages:
        #index = pages[4]
        print("Processing "+index)
        profiles = get_data(index)
            
        for profile in profiles:
            # if us_pa key exists get value and append to it, else add us_pa key with [] value and append to it
            profiles_dict.setdefault(f"{profile[20:22]}_{profile[23:25]}",[]).append(profile)
            
        for k, v in profiles_dict.items():
            with open(f"profiles/{k}_profiles", "a") as outfile:
                outfile.write("\n".join(v))
                outfile.write("\n")
        profiles_dict.clear()
        
  
    """
    for k, v in profiles.items():
        with open(f"profiles/{k}_profiles", "a") as outfile:
            outfile.write("\n".join(v))
    """