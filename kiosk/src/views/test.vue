<template>
    <LoginForm :register="false" actionButtonText="Login" />
</template>

<script>
    import LoginForm from '@/components/LoginForm.vue'

    export default {
        name: 'HomePage',
        components: {
            LoginForm
        },
        mounted() {

        }
    }
</script>
