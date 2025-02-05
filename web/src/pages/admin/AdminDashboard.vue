<template>
  <div class="flex h-screen bg-gray-100">
    <aside class="hidden md:block">
      <Sidebar />
    </aside>

    <!-- Sidebar untuk mobile (drawer) -->
    <transition name="slide">
      <div v-if="showSidebar" class="fixed inset-0 z-50 flex">
        <!-- Overlay -->
        <div
          class="absolute inset-0"
          @click="toggleSidebar"
        ></div>
        <!-- Sidebar Mobile -->
        <div class="relative bg-white w-64 shadow-lg">
          <Sidebar />
        </div>
      </div>
    </transition>

    <!-- Konten Utama -->
    <div class="flex flex-col flex-1">
      <!-- Header Mobile -->
      <header class="md:hidden flex items-center justify-between bg-white shadow p-4">
        <button @click="toggleSidebar" class="text-gray-600 focus:outline-none">
          <!-- Icon Hamburger -->
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h16M4 18h16"
            />
          </svg>
        </button>
        <h1 class="text-lg font-bold text-blue-600">IP-CORP Admin</h1>
        <!-- Placeholder agar header seimbang -->
        <div class="w-6"></div>
      </header>

      <!-- Main Content -->
      <main class="flex-1 p-8 overflow-auto">
        <div class="flex justify-between items-center">
          <div>
            <h2 class="text-xl font-semibold">Selamat Datang, Admin!</h2>
            <p class="text-gray-500">Home > Dashboard</p>
          </div>
          <!-- Foto Profil -->
          <img
            src="/img/background/background.png"
            alt="User Profile"
            class="w-10 h-10 rounded-full border-2 border-gray-300 shadow-md cursor-pointer"
          />
        </div>

        <!-- Dashboard Stats -->
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-6 mt-8">
          <div class="bg-white p-6 shadow-md rounded-md flex items-center space-x-4">
            <div class="text-blue-600 text-3xl">🅱️</div>
            <div>
              <h3 class="text-2xl font-semibold">237</h3>
              <p class="text-gray-500">Blog</p>
            </div>
          </div>

          <div class="bg-white p-6 shadow-md rounded-md flex items-center space-x-4">
            <div class="text-blue-600 text-3xl">🍔</div>
            <div>
              <h3 class="text-2xl font-semibold">177</h3>
              <p class="text-gray-500">Makanan</p>
            </div>
          </div>

          <div class="bg-white p-6 shadow-md rounded-md flex items-center space-x-4">
            <div class="text-blue-600 text-3xl">📖</div>
            <div>
              <h3 class="text-2xl font-semibold">31</h3>
              <p class="text-gray-500">Resep</p>
            </div>
          </div>

          <div class="bg-white p-6 shadow-md rounded-md flex items-center space-x-4">
            <div class="text-blue-600 text-3xl">🏋️‍♂️</div>
            <div>
              <h3 class="text-2xl font-semibold">31</h3>
              <p class="text-gray-500">Kebugaran</p>
            </div>
          </div>

          <div class="bg-white p-6 shadow-md rounded-md flex items-center space-x-4">
            <div class="text-blue-600 text-3xl">👤</div>
            <div>
              <h3 class="text-2xl font-semibold">1,653</h3>
              <p class="text-gray-500">User</p>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script>
import Sidebar from "./components/Sidebar.vue";

export default {
  name: "DashboardLayout",
  components: {
    Sidebar,
  },
  data() {
    return {
      showSidebar: false,
    };
  },
  methods: {
    toggleSidebar() {
      this.showSidebar = !this.showSidebar;
    },
  },
};
</script>

<style>
.slide-enter-active,
.slide-leave-active {
  transition: transform 0.3s ease;
}
.slide-enter-from,
.slide-leave-to {
  transform: translateX(-100%);
}
</style>
