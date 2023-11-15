<script>
	import { Router, Route, navigate } from "svelte-routing";

	import Menu from "../components/Menu.svelte";
	import Table from "../components/Table.svelte";
	import { DateInput } from "date-picker-svelte";

	//import Modal from './Modal.svelte';
	//let showModal = false;
	import { Button, Modal, FormGroup, Input, Label, Row, Col, Badge } from "sveltestrap";


	async function getMembers() {
		let resp = await fetch(
			"/api/members"
		).then((x) => x.json());
		//data.set(resp.members)

		if (resp.members){
			return resp.members
		};
	}
	let promise = getMembers();
	//let date;
	//data.set(resp.members)
	
    let open = false;
    const toggle = () => (open = !open);
	

    async function onSubmit(e) {
		const formData = new FormData(e.target);

		const data = {};
		for (let field of formData) {
			const [key, value] = field;
			data[key] = value;
		}

		const res = await fetch('/api/members', {
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
	$: dob_date_iso = dob_date.toISOString()
	$: id_date_iso = id_date.toISOString()
</script>

<Menu />
{#await promise}
	<p>Loading...</p>
{:then members}
<Row class="mx-auto my-2">
    <Col >
        <Button  on:click={toggle}>Add Member <Badge>+</Badge></Button>
    </Col>
    <Col sm="3">		
        <Input type="text" bind:value={$filterValue} placeholder="Search rows..." />
    </Col>
</Row>
{#if members}
<Table data={members} bind:filterValue columns={Object.keys(members[0])}>
		
</Table>
{:else}
<p>No members found</p>
{/if}

{/await}
<Modal body  header="Registration Form" backdrop="static" isOpen={open} {toggle} scrollable>
		
    <form on:submit|preventDefault={onSubmit}>
    <div class="container">  
        <FormGroup floating label="First Name">
            <Input name="first_name" placeholder="First Name" required/>
        </FormGroup>
        <FormGroup floating label="Middle Name">
            <Input name="middle_name" placeholder="Middle Name" required/>
        </FormGroup>
        <FormGroup floating label="Last Name">
            <Input name="last_name" placeholder="Last Name" required/>
        </FormGroup>
        <Label class=""><b>Gender</b></Label>
        <div class="form-check">
            <input class="form-check-input" type="radio" name="gender" id="radio1" value="Male" required>
            <label class="form-check-label" for="radio1">
              Male
            </label>
          </div>
          <div class="form-check">
            <input class="form-check-input" type="radio" name="gender" id="radio2" value="Female" required>
            <label class="form-check-label" for="radio2">
              Female
            </label>
        </div>
        <br>
		<FormGroup >
            <Label>Date of Birth</Label>
			<input name="dob" bind:value={dob_date_iso} style="display:none;">
            <DateInput bind:value={dob_date} format="dd-MM-yyyy" min={new Date("01-01-1900")} closeOnSelection=true />
        </FormGroup>
        <FormGroup floating label="ID Number">
            <Input name="id_no" placeholder="ID Number" required/>
        </FormGroup>
		<FormGroup >
            <Label>ID Issue date</Label>
			<input name="pp" bind:value={id_date_iso} style="display:none;">
            <DateInput bind:value={id_date} format="dd-MM-yyyy" min={new Date("01-01-1900")} closeOnSelection=true />
        </FormGroup>
		<FormGroup floating label="ID Issue Place">
            <Input name="id_issue_place" placeholder="ID Issue Place" required/>
        </FormGroup>		
		
        <FormGroup floating label="Phone">
            <Input name="phone" placeholder="Phone" required/>
        </FormGroup>		
        <FormGroup floating label="Email">
            <Input name="email" placeholder="Email" required/>
        </FormGroup>	
      
      <br>
      <div class="text-center">
      <Button type="submit" >Submit</Button></div>
    </div>
    
</form>

</Modal>
<!--style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}
	table {
		border-spacing: 0;
		border-top: 1px solid black;
		border-left: 1px solid black;
	}
	th,
	td {
		border-bottom: 1px solid black;
		border-right: 1px solid black;
		padding: 0.5rem;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style-->
