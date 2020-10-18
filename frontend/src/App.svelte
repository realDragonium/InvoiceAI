<script>
import Signup from './components/Signup.svelte'
import { Router, Link, Route } from "svelte-routing";
import Welcome from './components/Welcome.svelte';
</script>

<main>
  <Signup/>
  <div class="container">
    <Router>
  <nav>
    <Link to="/welcome">Home</Link>
    <Link to="about">About</Link>
    <Link to="blog">Blog</Link>
  </nav>
  

  <div>
    <Route path="/welcome" component="{Welcome}" />
  </div>
</Router>
  </div>
 

 
</main>

