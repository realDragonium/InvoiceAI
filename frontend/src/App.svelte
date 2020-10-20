<script>
    import Router from 'svelte-spa-router'
    import Auth from './components/auth/Auth.svelte';
    import Welcome from './components/Welcome.svelte';
    import Home from './components/Home.svelte';

    const routes = {
        '/home': Home,
        '/auth': Auth,
        '/welcome': Welcome,
    }

</script>


<main>
    <a href="#/home" >home</a>
    <a href="#/welcome">welcome</a>
    <a href="#/auth">Auth</a>
    <Router {routes}/>
</main>

