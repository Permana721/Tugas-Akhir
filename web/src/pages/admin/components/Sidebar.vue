<template>
  <aside class="w-64 h-full bg-white shadow-lg p-6">
    <RouterLink to="/admin-pg">
      <h1 class="text-2xl mt-2 font-bold text-blue-600">IP-CORP Admin</h1>
    </RouterLink>

    <nav class="mt-6">
      <!-- Section Main -->
      <h2 class="text-gray-500 uppercase text-sm font-semibold">Main</h2>
      <ul class="mt-2 space-y-2">
        <li class=" duration-300 hover:bg-blue-100 rounded-md">
          <router-link
            to="/admin-pg"
            class="flex items-center p-2 rounded-md"
            :class="{ 'bg-blue-100 text-blue-600 font-semibold': $route.path === '/admin-pg' }"
          >
            Dashboard
          </router-link>
        </li>
      </ul>

      <!-- Section Content -->
      <h2 class="mt-6 text-gray-500 uppercase text-sm font-semibold">Content</h2>
      <ul class="mt-2 space-y-2">
        <li class="duration-300 hover:bg-blue-100 rounded-md">
          <router-link
            to="/admin-pg/blog"
            class="p-2 block rounded-md"
            :class="{ 'bg-blue-100 text-blue-600 font-semibold': $route.path === '/admin-pg/blog' }"
          >
            Blog
          </router-link>
        </li>
        <li class="duration-300 hover:bg-blue-100 rounded-md">
          <router-link
            to="/admin-pg/food"
            class="p-2 block rounded-md"
            :class="{ 'bg-blue-100 text-blue-600 font-semibold': $route.path === '/admin-pg/food' }"
          >
            Makanan
          </router-link>
        </li>
        <li class="duration-300 hover:bg-blue-100 rounded-md">
          <router-link
            to="/admin-pg/recipe"
            class="p-2 block rounded-md"
            :class="{ 'bg-blue-100 text-blue-600 font-semibold': $route.path === '/admin-pg/recipe' }"
          >
            Resep
          </router-link>
        </li>
      </ul>

      <!-- Section Account -->
      <h2 class="mt-6 text-gray-500 uppercase text-sm font-semibold">Account</h2>
      <ul class="mt-2 space-y-2">
        <li class="duration-300 hover:bg-blue-100 rounded-md">
          <router-link
            to="/admin-pg/user"
            class="p-2 block rounded-md"
            :class="{ 'bg-blue-100 text-blue-600 font-semibold': $route.path === '/admin-pg/user' }"
          >
            User
          </router-link>
        </li>
      </ul>
    </nav>
  </aside>
</template>

<script>
import { RouterLink } from 'vue-router';

export default {
  name: "AdminSidebar",
};
</script>
