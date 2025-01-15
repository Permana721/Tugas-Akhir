<template>
    <footer class=" bg-white shadow-2xl text-gray-400 py-8 z-50">
        <div class="max-w-7xl px-4 sm:px-6 lg:px-8 mx-auto flex flex-col md:flex-row justify-between items-center">
            <RouterLink to="/" class="flex items-center space-x-2 mb-4 md:mb-0">
                <img src="/img/icon/logo.png" alt="My Diet Coach Logo" class="w-40" />
            </RouterLink>
    
            <nav class="flex space-x-6 text-black text-sm">
                <RouterLink to="/" class=" hover:text-blue-600 duration-300">Home</RouterLink>
                <RouterLink to="" class="hover:text-blue-600 duration-300">About</RouterLink>
                <RouterLink to="" class="hover:text-blue-600 duration-300">Contact Us</RouterLink>
                </nav>
        </div>
    
        <div class="mt-8 text-center border-t border-gray-700 pt-4">
            <p class="text-sm">
            © {{ currentYear }} FUFUFAFA - All rights reserved
            </p>
        </div>
    </footer>
</template>

<script>
    export default {
        name: "Footer",
        data() {
            return {
            currentYear: new Date().getFullYear(),
            };
        },
    };
</script>
