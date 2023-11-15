<script>
    import { Link, navigate } from 'svelte-routing';
    import { token } from '../stores.js';
    import { Input, Button, Container, Col, Row } from 'sveltestrap';

    async function login(e) {
		const formData = new FormData(e.target);

		const data = {};
		for (let field of formData) {
			const [key, value] = field;
			data[key] = value;
		}

		const res = await fetch('/api/users/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Accept': 'application/json',
			},
			body: JSON.stringify(data)
		})
        if (res.ok) {
            const json = await res.json()
            const result = JSON.stringify(json)
            console.log(data)
            console.log(result)
            localStorage.setItem('token', '1');
            token.set(localStorage.getItem('token'));

            navigate('/members', { replace: false });
        }
        else {
            navigate('/denied')
        }
		console.log(res)

    };
</script>

<!--h1>👋 Welcome to the login page</h1>
<h2>
    <Link to="members">Go to Dashboard</Link>
</h2>
<span>or</span-->
<Container class="">
    <div class="d-inline-flex position-absolute top-50 start-50 translate-middle rounded-4 border border-dark border-3 mx-auto bg-info align-items-center" style="height:25rem;width:25rem;" >

        <Col class="pt-auto align-self-center" sm={{ size: 8, offset: 2 }}>
            <form on:submit|preventDefault={login}>
                <Input class="border border-dark border-3" name="username" placeholder="Username" required/>
                <br>
                <Input class="border border-dark border-3" name="password" placeholder="Password" type="password" required/>
                <br>
                <br>
                <br>
                <Button class="position-absolute start-50 top-75 bg-light text-dark translate-middle border border-dark border-3" on:click={login}><b>Login</b></Button>
            </form>
        </Col>
    </div>
</Container>