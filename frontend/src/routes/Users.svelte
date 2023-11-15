<script>
	import { Router, Route, navigate } from "svelte-routing";

	import Menu from "../components/Menu.svelte";
	import Table from "../components/Table.svelte";
	import { DateInput } from "date-picker-svelte";

	//import Modal from './Modal.svelte';
	//let showModal = false;
	import { Button, Modal, FormGroup, Input, Label, Row, Col, Badge } from "sveltestrap";


	async function getUsers() {
		let resp = await fetch(
			"/api/users"
		).then((x) => x.json());
		//data.set(resp.users)

		return resp.users;
	}
	let promise = getUsers();
	//let date;
	//data.set(resp.users)
	
    let open = false;
    const toggle = () => (open = !open);
	

    async function onSubmit(e) {
		const formData = new FormData(e.target);

		const data = {};
		for (let field of formData) {
			const [key, value] = field;
			data[key] = value;
		}

		const res = await fetch('/api/users', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Accept': 'application/json',
			},
			body: JSON.stringify(data)
		})
		toggle()
		
		const json = await res.json()
		const result = JSON.stringify(json)
		console.log(data)
		console.log(result)
	}

	let filterValue;
	let dob_date =new Date();
	let id_date =new Date();
	let enabled = true;
	$: dob_date_iso = dob_date.toISOString()
	$: id_date_iso = id_date.toISOString()
</script>

<Menu />
{#await promise}
	<p>Loading...</p>
{:then users}
<Row class="mx-auto my-2">
    <Col >
        <Button  on:click={toggle}>Add User <Badge>+</Badge></Button>
    </Col>
    <Col sm="3">		
        <Input type="text" bind:value={$filterValue} placeholder="Search rows..." />
    </Col>
</Row>
	<Table data={users} bind:filterValue columns={Object.keys(users[0])}>
		
	</Table>
{/await}
<Modal body  header="Registration Form" backdrop="static" isOpen={open} {toggle} scrollable>
		
    <form on:submit|preventDefault={onSubmit}>
    <div class="container">  
        <FormGroup floating label="Name">
            <Input name="name" placeholder="Name" required/>
        </FormGroup>
        <FormGroup floating label="Username">
            <Input name="username" placeholder="Username" required/>
        </FormGroup>
        <FormGroup floating label="Password">
            <Input name="password" type="password" placeholder="Last Name" required/>
        </FormGroup>
		<FormGroup floating label="Role">
            <Input name="role" type="select" required>
				<option>Standard</option>
				<option>Admin</option>
			</Input>
        </FormGroup>
		<FormGroup floating>

        <Input class="form-check-input" type="checkbox" name="enabled" bind:value={enabled} bind:checked={enabled} required/>
        </FormGroup>
      <br>
      <div class="text-center">
      <Button type="submit" >Submit</Button></div>
    </div>
    
</form>

</Modal>