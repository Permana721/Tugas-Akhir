<template>
    <Navbar />
    <div class="flex flex-col">
        <div class="flex-1 pt-28">
            <h2 class="text-slate-600">Blog</h2>
            <p class="text-red-600">Welcome to the blog!</p>
        </div>
    </div>
    <Footer />
</template>

<script>
import Navbar from '../../../components/Navbar.vue';
import Footer from '../../../components/Footer.vue';

export default {
    components: {
        Navbar,
        Footer,
    },
};
</script>