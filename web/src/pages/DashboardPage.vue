<template>
    <div class="flex flex-col">
        <Navbar />
        <div class="flex-1 pt-28">
            <h2 class="text-slate-600">Dashboard</h2>
            <p class="text-red-600">Welcome to the dashboard!</p>
        </div>
    </div>
</template>

<script>
import Navbar from '../components/Navbar.vue';

export default {
components: { Navbar },
};
</script>