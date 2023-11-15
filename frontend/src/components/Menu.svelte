<script>
    import {  
        Collapse,
        Navbar,
        NavbarToggler,
        NavbarBrand,
        Nav,
        NavItem,
        NavLink,
        Image,
        Row,
        Col
    } from 'sveltestrap'
    import { navigate } from 'svelte-routing'
    let isNavOpen = false;

    function handleUpdate(event) {
        isNavOpen = event.detail.isNavOpen;
    }

    import { Link } from 'svelte-routing';
    import { token } from '../stores.js';

    const logout = () => {
        localStorage.clear();
        token.set(localStorage.getItem('token'));
        navigate('/', { replace: true });
    }
</script>
<Navbar color="light" light class="py-0" expand="md">
    <NavbarBrand href="/">
        <Row class="align-items-center">
          <Col><Image height="70vh" width="70vw" src='./assets/logo.png'/></Col>
          <Col class="text-center"><h5><b>Ansaar Community</b></h5></Col>
        </Row>
    </NavbarBrand>
    <NavbarToggler on:click={() => (isNavOpen = !isNavOpen)} />
    <Collapse {isNavOpen} navbar expand="md" on:update={handleUpdate}>
      <Nav class="ms-auto" navbar>
        <NavItem>
          <Link class="nav-link" to="/members">Members</Link>
        </NavItem>
        <NavItem>
            <NavLink  href="#">Payments</NavLink>
          </NavItem>
          <NavItem>
            <NavLink  href="#">Requests</NavLink>
          </NavItem>
          <NavItem>
            <Link class="nav-link" to="/users">Users</Link>
          </NavItem>
        <NavItem>
          <NavLink on:click={logout} >Logout</NavLink>
        </NavItem>

      </Nav>
    </Collapse>
  </Navbar>