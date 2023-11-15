<script>
    import { Router, Route } from 'svelte-routing';
    import ProtectedRoute from './ProtectedRoute.svelte';

    import Login from './routes/Login.svelte';
    import Members from './routes/Members.svelte';
    import Users from './routes/Users.svelte';
    import Denied from './routes/Denied.svelte';
</script>
<svelte:head>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.1/dist/css/bootstrap.min.css">
  <link rel="stylesheet" type="text/css" href="custom.css">
</svelte:head>
<Router>
    <Route path="/" component={Login} />
    <Route path="/denied" component={Denied} />
    <ProtectedRoute path="/members" component={Members} />
    <ProtectedRoute path="/users" component={Users} />
</Router>