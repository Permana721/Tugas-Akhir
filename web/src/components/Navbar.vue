<template>
    <nav class="bg-white shadow-lg fixed w-full z-50">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex justify-between h-20 items-center">
                <div class="flex items-center">
                    <img src="https://via.placeholder.com/40" alt="Logo" class="h-8 w-8 mr-2">
                    <span class="text-xl font-extrabold">YNKTS</span>
                </div>
        
                <div class="md:hidden">
                    <button
                    @click="toggleMenu"
                    class="focus:outline-none transition-transform transform"
                    :class="{ 'rotate-90': menuOpen }"
                    >
                    <svg
                        v-if="!menuOpen"
                        class="w-8 h-8 text-purple-600"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
                    </svg>
                    <svg
                        v-else
                        class="w-8 h-8 text-purple-600"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                    </button>
                </div>
        
                <div class="hidden md:flex md:items-center space-x-4">
                    <a href="#" class="text-gray-600 hover:text-purple-600 font-medium p-2">Blog</a>
                    <a href="#" class="text-gray-600 hover:text-purple-600 font-medium p-2">Diet</a>
                    <a href="#" class="text-gray-600 hover:text-purple-600 font-medium p-2">Rencana</a>
                    <a href="#" class="text-gray-600 hover:text-purple-600 font-medium p-2">Makan</a>
                    <a href="#" class="text-gray-600 hover:text-purple-600 font-medium p-2">Resep</a>
                    <a href="#" class="text-gray-600 hover:text-purple-600 font-medium p-2">Produk</a>
                </div>
        
                <div class="hidden md:block">
                    <a href="#" class="text-purple-600 px-4 py-2 rounded-md border border-purple-700">
                    Masuk
                    </a>
                </div>
            </div>
        </div>
        <transition name="slide-down">
            <div
                v-if="menuOpen"
                class="bg-purple-600 text-white fixed top-20 left-0 w-full h-1/2 z-40"
            >
                <div class="flex flex-col items-center justify-center h-full space-y-4">
                    <a href="#" class="text-xl font-medium">Blog</a>
                    <a href="#" class="text-xl font-medium">Diet</a>
                    <a href="#" class="text-xl font-medium">Rencana</a>
                    <a href="#" class="text-xl font-medium">Makan</a>
                    <a href="#" class="text-xl font-medium">Resep</a>
                    <a href="#" class="text-xl font-medium">Produk</a>
                    <a href="#" class="bg-white text-purple-600 px-4 py-2 rounded-md font-bold border border-white">
                        Masuk
                    </a>
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