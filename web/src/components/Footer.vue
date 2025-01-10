<template>
    <footer class="bg-gray-100 text-gray-400 py-8">
        <div class="max-w-6xl mx-auto px-4 flex flex-col md:flex-row justify-between items-center">
            <div class="flex items-center space-x-2 mb-4 md:mb-0">
                <img src="/img/icon/logo.png" alt="My Diet Coach Logo" class=" w-36" />
            </div>
    
            <nav class="flex space-x-6 text-black text-sm">
                <RouterLink to="/" class="hover:text-white duration-300">Home</RouterLink>
                <RouterLink to="/about" class="hover:text-white duration-300">About</RouterLink>
                <RouterLink to="/contact" class="hover:text-white duration-300">Contact Us</RouterLink>
                </nav>
        </div>
    
        <div class="mt-8 text-center border-t border-gray-700 pt-4">
            <p class="text-sm">
            © 2025 FUFUFAFA - All rights reserved
            </p>
        </div>
    </footer>
</template>

<script>
    export default {
        name: "Footer",
    };
</script>
