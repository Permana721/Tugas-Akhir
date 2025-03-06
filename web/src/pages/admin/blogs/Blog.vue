<template>
  <div class="flex flex-col md:flex-row min-h-screen bg-gray-100">
    <!-- Sidebar -->
    <aside
      class="inset-y-0 left-0 w-64 bg-white shadow-lg transition-transform transform md:translate-x-0 z-50
            fixed md:static"
      :class="{ '-translate-x-full': !isSidebarOpen }"
    >
      <Sidebar />
    </aside>

    <!-- Overlay untuk menutup sidebar di mobile -->
    <div
      v-if="isSidebarOpen"
      class="fixed inset-0 bg-black bg-opacity-50 md:hidden"
      @click="isSidebarOpen = false"
    ></div>

    <!-- Konten utama -->
    <div class="flex-1 p-4 md:p-6">
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-xl md:text-2xl font-bold text-blue-600">Kelola Blog</h2>
        <button 
          @click="isSidebarOpen = !isSidebarOpen" 
          class="md:hidden text-blue-600 transition-transform duration-300 ease-in-out"
          :class="{ 'rotate-90': isSidebarOpen }"
        >
          ☰
        </button>
      </div>

      <div class="bg-white p-4 shadow-md rounded-md overflow-x-auto">
        <div class="flex flex-col md:flex-row md:justify-between md:items-center gap-2 md:gap-0 mb-4">
          <!-- Ganti tombol dengan router-link -->
          <router-link
            to="/admin-pg/blog/add-blog"
            class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 w-full md:w-auto text-center"
          >
            + Tambah Blog
          </router-link>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Cari blog..."
            class="w-full md:w-56 p-2 border rounded-md"
          />
        </div>

        <!-- Tabel -->
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-blue-50">
              <th class="p-4 border-b-2">Judul</th>
              <th class="p-4 border-b-2">Deskripsi</th>
              <th class="p-4 border-b-2">Kategori</th>
              <th class="p-4 border-b-2 hidden md:table-cell">Penulis</th>
              <th class="p-4 border-b-2">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="blog in filteredBlogs" :key="blog.id" class="hover:bg-gray-50">
              <td class="p-4 border-b">{{ blog.title }}</td>
              <td class="p-4 border-b">{{ blog.description }}</td>
              <td class="p-4 border-b">{{ blog.category }}</td>
              <td class="p-4 border-b hidden md:table-cell">{{ 'Admin' }}</td>
              <td class="p-4 border-b flex flex-col md:flex-row space-y-2 md:space-y-0 md:space-x-2">
                <button
                  class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 w-full md:w-auto"
                  @click="editBlog(blog)"
                >
                  Edit
                </button>
                <button
                  class="px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 w-full md:w-auto"
                  @click="deleteBlog(blog.id)"
                >
                  Hapus
                </button>
              </td>
            </tr>
            <tr v-if="filteredBlogs.length === 0">
              <td colspan="3" class="text-center p-4 text-gray-500">
                {{ isLoading ? "Memuat data..." : "Tidak ada blog ditemukan." }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Sidebar from "../components/Sidebar.vue";

export default {
  components: {
    Sidebar,
  },
  data() {
    return {
      isSidebarOpen: false,
      searchQuery: "",
      blogs: [],
      isLoading: false, 
    };
  },
  computed: {
    filteredBlogs() {
      return this.blogs.filter((blog) =>
        blog.title.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    },
  },
  methods: {
    async fetchBlogs() {
      this.isLoading = true;
      try {
        const response = await axios.get("http://localhost:8000/blog");
        this.blogs = response.data;
      } catch (error) {
        console.error("Gagal mengambil data blog:", error);
        alert("Gagal memuat data blog. Cek koneksi API.");
      } finally {
        this.isLoading = false;
      }
    },
    editBlog(blog) {
      this.$router.push({
        path: `/admin-pg/blog/edit-blog/${blog.id}`,
        state: { blog }, 
      });
    },
    async deleteBlog(id) {
      if (!confirm("Apakah Anda yakin ingin menghapus blog ini?")) return;
      try {
        await axios.delete(`http://localhost:8000/blog/delete/${id}`);
        this.blogs = this.blogs.filter((blog) => blog.id !== id);
        alert("Blog berhasil dihapus.");
      } catch (error) {
        console.error("Gagal menghapus blog:", error);
        alert("Gagal menghapus blog. Cek koneksi API.");
      }
    },
  },
  mounted() {
    this.fetchBlogs();
  },
};
</script>