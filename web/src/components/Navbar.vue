<template>
    <nav class="bg-white shadow-md fixed w-full z-50">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex justify-between h-20 items-center">
                <div class="flex items-center">
                    <RouterLink to="/">
                        <img src="/img/icon/logos.png" alt="Logo" class="w-40">
                    </RouterLink>
                </div>
                <div class="md:hidden">
                    <button
                    @click="toggleMenu"
                    class="focus:outline-none transition-transform transform"
                    :class="{ 'rotate-90': menuOpen }"
                    >
                    <svg
                        v-if="!menuOpen"
                        class="w-8 h-8 text-blue-600"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
                    </svg>
                    <svg
                        v-else
                        class="w-8 h-8 text-blue-600"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                    </button>
                </div>
        
                <div class="hidden md:flex md:items-center justify-center text-center mr-16 space-x-8">
                    <RouterLink to="/dashboard" :class="{'text-blue-600 font-semibold': $route.path === '/dashboard', 'text-gray-600 hover:text-blue-600 font-medium': $route.path !== '/dashboard'}" class="p-2">
                    Dashboard
                    </RouterLink>
                    <RouterLink to="/blog" :class="{'text-blue-600 font-semibold': $route.path === '/blog', 'text-gray-600 hover:text-blue-600 font-medium': $route.path !== '/blog'}" class="p-2">
                    Blog
                    </RouterLink>
                    <RouterLink to="/food" :class="{'text-blue-600 font-semibold': $route.path === '/food', 'text-gray-600 hover:text-blue-600 font-medium': $route.path !== '/food'}" class="p-2">
                    Makanan
                    </RouterLink>
                    <RouterLink to="/sport" :class="{'text-blue-600 font-semibold': $route.path === '/sport', 'text-gray-600 hover:text-blue-600 font-medium': $route.path !== '/sport'}" class="p-2">
                    Kebugaran
                    </RouterLink>
                    <RouterLink to="/recipe" :class="{'text-blue-600 font-semibold': $route.path === '/recipe', 'text-gray-600 hover:text-blue-600 font-medium': $route.path !== '/recipe'}" class="p-2">
                    Resep
                    </RouterLink>
                </div>
        
                <div class="hidden md:block">
                    <RouterLink to="/login" class="text-blue-600 px-4 py-2 rounded-md border border-blue-700">
                    Masuk
                    </RouterLink>
                </div>
            </div>
        </div>
        <transition name="slide-down">
            <div
                v-if="menuOpen"
                class="bg-blue-600 fixed text-white top-20 left-0 w-full h-80 z-40"
            >
                <div class="flex flex-col items-center justify-center h-full space-y-4">
                    <RouterLink to="/dashboard" :class="{'text-black text-xl font-medium': $route.path === '/dashboard', 'text-white text-xl font-medium': $route.path !== '/dashboard'}">Dashboard</RouterLink>
                    <RouterLink to="/blog" :class="{'text-black text-xl font-medium': $route.path === '/blog', 'text-white text-xl font-medium': $route.path !== '/blog'}">Blog</RouterLink>
                    <RouterLink to="/food" :class="{'text-black text-xl font-medium': $route.path === '/food', 'text-white text-xl font-medium': $route.path !== '/food'}">Makanan</RouterLink>
                    <RouterLink to="/recipe" :class="{'text-black text-xl font-medium': $route.path === '/recipe', 'text-white text-xl font-medium': $route.path !== '/recipe'}">Resep</RouterLink>
                    <RouterLink to="/sport" :class="{'text-black text-xl font-medium': $route.path === '/sport', 'text-white text-xl font-medium': $route.path !== '/sport'}">Kebugaran</RouterLink>
                    <RouterLink
                        to="/login"
                        class="bg-white text-blue-600 px-4 py-2 rounded-md font-bold border border-white"
                    >
                        Masuk
                    </RouterLink>
                </div>
            </div>
        </transition>
    </nav>
</template>

<script>
export default {
data() {
    return {
    menuOpen: false, 
    };
},
methods: {
    toggleMenu() {
    this.menuOpen = !this.menuOpen;
    },
},
};
</script>

<style>
.slide-down-enter-active {
animation: slide-down 0.5s ease-out forwards;
}
.slide-down-leave-active {
animation: slide-up 0.5s ease-in forwards;
}

@keyframes slide-down {
0% {
    opacity: 0;
    transform: translateY(-20%);
}
100% {
    opacity: 1;
    transform: translateY(0);
}
}

@keyframes slide-up {
0% {
    opacity: 1;
    transform: translateY(0);
}
100% {
    opacity: 0;
    transform: translateY(-100%);
}
}
</style>