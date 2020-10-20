<script>
    import Signup from './Signup.svelte';
    import Signin from './Signin.svelte';
    import { fade } from 'svelte/transition'

    const options = {
        'signin': {
            component: Signin
        },
        'signup': {
            component: Signup
        }
    };

    let viewportComponent = options.signin;
    let currentSelected = viewportComponent;

    function handleGoto(event) {
        currentSelected = options[event.detail.component];
    }

    function updateViewportComponent() {
        viewportComponent = currentSelected;
    }

</script>

<style>
    .login-dark {
        height: 1000px;
        background: rgba(0, 0, 0, .65) url(../images/ai.jpg);
        background-blend-mode: darken;
        background-size: cover;
        position: relative;
        display: grid;
        place-items: center;
    }

    .login-dark .form {
        background-color: #1e2833;
        padding: 40px;
        border-radius: 4px;
        color: #fff;
        box-shadow: 3px 3px 4px rgba(0, 0, 0, 0.2);
        overflow: hidden;
    }

</style>

<div class="login-dark">
    <div class="form">
        {#if viewportComponent == currentSelected}
            <div on:outroend={updateViewportComponent} transition:fade>
                <svelte:component this={viewportComponent.component} on:message={handleGoto}/>
            </div>
        {/if}
    </div>
</div>



