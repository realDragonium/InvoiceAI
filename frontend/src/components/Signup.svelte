<script>
    import {navigate} from "svelte-routing";

    let username;
    let email;
    let password;
    let hide = false
    const handleSubmit = async () => {
        try {
            const data = {username, email, password}
            const res = await fetch('/api/v1/register', {
                method: 'POST',
                body: JSON.stringify(data),
                credentials: 'include',
                headers: {'Content-Type': 'application/json'},

            })

            const returnData = await res.json()
            username = ''
            email = ''
            password = ''
            hide = true
        } catch (error) {
            console.log(error)

        }
        navigate("/welcome");


    }
</script>

<style>
    .login-dark {
        height: 1000px;
        background: rgba(0, 0, 0, .65) url(../images/ai.jpg);
        background-blend-mode: darken;
        background-size: cover;
        position: relative;
    }

    .login-dark form {
        max-width: 320px;
        width: 90%;
        background-color: #1e2833;
        padding: 40px;
        border-radius: 4px;
        transform: translate(-50%, -50%);
        position: absolute;
        top: 50%;
        left: 50%;
        color: #fff;
        box-shadow: 3px 3px 4px rgba(0, 0, 0, 0.2);
    }

    .login-dark .illustration {
        text-align: center;
        padding: 15px 0 20px;
        font-size: 100px;
        color: #2980ef;
    }

    .login-dark form .form-control {
        background: none;
        border: none;
        border-bottom: 1px solid #434a52;
        border-radius: 0;
        box-shadow: none;
        outline: none;
        color: inherit;
    }

    .login-dark form .btn-primary {
        background: #214a80;
        border: none;
        border-radius: 4px;
        padding: 11px;
        box-shadow: none;
        margin-top: 26px;
        text-shadow: none;
        outline: none;
    }

    .login-dark form .btn-primary:hover, .login-dark form .btn-primary:active {
        background: #214a80;
        outline: none;
    }


    .login-dark form .btn-primary:active {
        transform: translateY(1px);
    }


</style>
{#if !hide}
    <div class="login-dark">
        <form on:submit|preventDefault="{handleSubmit}">
            <h2 class="sr-only">Signup Form</h2>
            <div class="illustration"><i class="icon ion-ios-locked-outline"></i></div>
            <div class="form-group"><input class="form-control" type="username" name="username" bind:value="{username}"
                                           placeholder="Username"/></div>
            <div class="form-group"><input class="form-control" type="email" name="email" bind:value="{email}"
                                           placeholder="Email"/></div>
            <div class="form-group"><input class="form-control" type="password" name="password" bind:value="{password}"
                                           placeholder="Password"/></div>
            <div class="form-group">
                <button class="btn btn-primary btn-block" type="submit">Register</button>
            </div>
        </form>
    </div>
{/if}


