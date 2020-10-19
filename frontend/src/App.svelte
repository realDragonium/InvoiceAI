<script>
    import Signup from './components/Signup.svelte';
    import Welcome from './components/Welcome.svelte';
    import Home from './components/Home.svelte';

    import Router from 'svelte-spa-router';


    const routes = {
        '/signup': Signup,
        '/welcome': Welcome,
        '/home': Home,

    }

</script>


<body>
  <Router {routes}/>
</body>


<!--<Signup/>-->
<!--<div class="container">-->
<!--  <Router>-->
<!--    <nav>-->
<!--      <Link to="/welcome">Home</Link>-->
<!--      <Link to="about">About</Link>-->
<!--      <Link to="blog">Blog</Link>-->
<!--    </nav>-->


<!--    <div>-->
<!--      <Route path="/welcome" component="{Welcome}" />-->
<!--    </div>-->
<!--  </Router>-->
<!--</div>-->
