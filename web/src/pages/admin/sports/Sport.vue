<template>
  <div class="flex flex-col md:flex-row min-h-screen bg-gray-100">
    <aside
      class="inset-y-0 left-0 w-64 bg-white shadow-lg transition-transform transform md:translate-x-0 z-50
            fixed md:static"
      :class="{ '-translate-x-full': !isSidebarOpen }"
    >
      <Sidebar />
    </aside>
    <div
      v-if="isSidebarOpen"
      class="fixed inset-0 bg-black bg-opacity-50 md:hidden"
      @click="isSidebarOpen = false"
    ></div>
    <div class="flex-1 p-4 md:p-6">
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-xl md:text-2xl font-bold text-blue-600">Kelola Olahraga</h2>
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
          <button
            class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 w-full md:w-auto"
            @click="showAddForm = true"
          >
            + Tambah Olahraga
          </button>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Cari olahraga..."
            class="w-full md:w-56 p-2 border rounded-md"
          />
        </div>
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-blue-50">
              <th class="p-4 border-b-2">Nama Olahraga</th>
              <th class="p-4 border-b-2 hidden md:table-cell">Deskripsi</th>
              <th class="p-4 border-b-2 hidden md:table-cell">Durasi</th>
              <th class="p-4 border-b-2 hidden md:table-cell">Jumlah Kalori</th>
              <th class="p-4 border-b-2">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="sport in filteredSports" :key="sport.id" class="hover:bg-gray-50">
              <td class="p-4 border-b">{{ sport.name }}</td>
              <td class="p-4 border-b hidden md:table-cell">{{ sport.description }}</td>
              <td class="p-4 border-b hidden md:table-cell">{{ sport.duration }}</td>
              <td class="p-4 border-b hidden md:table-cell">{{ sport.calories }}</td>
              <td class="p-4 border-b flex flex-col md:flex-row space-y-2 md:space-y-0 md:space-x-2">
                <button
                  class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 w-full md:w-auto"
                  @click="editSport(sport)"
                >
                  Edit
                </button>
                <button
                  class="px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 w-full md:w-auto"
                  @click="deleteSport(sport.id)"
                >
                  Hapus
                </button>
              </td>
            </tr>
            <tr v-if="filteredSports.length === 0">
              <td colspan="4" class="text-center p-4 text-gray-500">
                Tidak ada olahraga ditemukan.
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import Sidebar from "../components/Sidebar.vue";

export default {
  components: {
    Sidebar,
  },
  data() {
    return {
      isSidebarOpen: false,
      showAddForm: false,
      searchQuery: "",
      newSport: {
        name: "",
        description: "",
        duration: "",
        calories: "",
        fat: "",
      },
      sports: [
        { id: 1, name: "Jogging", description: "Latihan kardio ringan", duration: "30", calories: "200", fat: "5" },
        { id: 2, name: "Berenang", description: "Latihan seluruh tubuh", duration: "45", calories: "300", fat: "8" },
        { id: 3, name: "Yoga", description: "Latihan pernapasan dan fleksibilitas", duration: "60", calories: "150", fat: "3" },
      ],
    };
  },
  computed: {
    filteredSports() {
      return this.sports.filter((sport) =>
        sport.name.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    },
  },
  methods: {
    addSport() {
      if (!this.newSport.name || !this.newSport.description) {
        alert("Nama olahraga dan deskripsi harus diisi!");
        return;
      }
      this.sports.push({
        id: this.sports.length + 1,
        ...this.newSport,
      });
      this.newSport = {
        name: "",
        description: "",
        duration: "",
        calories: "",
        fat: "",
      };
      this.showAddForm = false;
    },
    editSport(sport) {
      alert(`Edit olahraga: ${sport.name}`);
    },
    deleteSport(id) {
      if (confirm("Apakah Anda yakin ingin menghapus olahraga ini?")) {
        this.sports = this.sports.filter((sport) => sport.id !== id);
      }
    },
  },
};
</script>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
}
th, td {
  padding: 12px;
}
</style>
