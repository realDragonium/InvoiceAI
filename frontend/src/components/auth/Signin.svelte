<script>
    import {push} from 'svelte-spa-router'
    import {Icon, TextField} from 'svelte-materialify/src';
    import Button from 'svelte-materialify/src/components/Button';
    import {mdiLockOutline, mdiLockOpenVariantOutline} from '@mdi/js';
    import {post} from "axios";
    import {createEventDispatcher} from 'svelte';

    import {user} from './stores.js'

    const dispatch = createEventDispatcher();
    const handleSubmit = async () => {
        post('/api/v1/login', $user).then(resp => {
            console.log(resp);
            push("/welcome");
        }).catch(err => {
            console.log(err)
        })
    }

    function goToSignUp() {
        dispatch('message', {
            component: 'signup'
        });
    }

</script>

<h2 class="sr-only">Login Form
    <Icon path={mdiLockOutline}/>
</h2>
<TextField type="username" name="username" bind:value="{$user.username}">Username</TextField>
<TextField type="email" name="email" bind:value="{$user.email}">Email</TextField>
<TextField type="password" name="password" bind:value="{$user.password}">Password</TextField>
<Button class="primary-color" block on:click={handleSubmit}>Login
    <Icon path={mdiLockOpenVariantOutline }/>
</Button>
<Button class="secondary-color" block on:click={goToSignUp}>
    Need an account?
</Button>
