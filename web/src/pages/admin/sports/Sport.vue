<template>
    <div class="flex h-screen">
        <aside class="hidden md:block">
            <Sidebar />
        </aside>
        <div class="flex-1 pt-28">
            <h2 class="text-slate-600">Sport</h2>
            <p class="text-red-600">Welcome to the Sport!</p>
        </div>
    </div>
</template>

<script>
import Sidebar from "../components/Sidebar.vue";

export default {
  name: "DashboardLayout",
  components: {
    Sidebar,
  },
}
</script>