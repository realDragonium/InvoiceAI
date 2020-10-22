
<script>
    import { Link } from "svelte-routing";
</script>




<style>
        @media (min-width: 800px) and (max-width: 850px) {
            .navbar:not(.top-nav-collapse) {
                background: #1C2331!important;
            }
        }
.navbar{background-color:rgba(0, 0, 0, 0.9)}@media only screen and (max-width:768px){.navbar{background-color:#090a0c}}
</style>



<nav class="navbar fixed-top navbar-expand-lg navbar-dark scrolling-navbar">
    <div class="container">
      <!-- Brand -->
      <a class="navbar-brand" href="https://mdbootstrap.com/material-design-for-bootstrap/" target="_blank">
        <strong>InvoiceAI</strong>
      </a>

      <!-- Collapse -->
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent"
        aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <!-- Links -->
      <div class="collapse navbar-collapse" id="navbarSupportedContent">

        <!-- Left -->
        <ul class="navbar-nav mr-auto">
          <li class="nav-item active">
            <Link to="/"><span class="nav-link">Home 
              <span class="sr-only">(current)</span>
            </span></Link>
          </li>
          <li class="nav-item">
           <Link to="/about"><span class="nav-link">About</span> </Link> 
          </li>
          <li class="nav-item">
            <Link to="/welcome"><span class="nav-link">Welcome</span></Link>
          </li>
        </ul>

        <!-- Right -->
        <ul class="navbar-nav nav-flex-icons">
          <li class="nav-item">
            <a href="https://github.com/realDragonium/InvoiceAI" class="nav-link border border-light rounded"
              target="_blank">
              <i class="fab fa-github mr-2"></i>InvoiceAI GitHub
            </a>
          </li>
        </ul>

      </div>

    </div>
  </nav>