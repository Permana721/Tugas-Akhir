<template>
    <div class="flex flex-col md:flex-row min-h-screen bg-gray-100">
        <aside
            class="inset-y-0 left-0 w-64 bg-white shadow-lg transition-transform transform md:translate-x-0 z-50
                fixed md:static"
            :class="{ '-translate-x-full': !isSidebarOpen }"
        >
            <Sidebar />
        </aside>

        <div v-if="isSidebarOpen" class="fixed inset-0 bg-black bg-opacity-50 md:hidden" @click="isSidebarOpen = false"></div>
        <div class="flex-1 p-4 md:p-6">
            <div class="flex justify-between items-center mb-4">
                <h2 class="text-xl md:text-2xl font-bold text-blue-600">Edit Blog</h2>
                <button 
                    @click="isSidebarOpen = !isSidebarOpen" 
                    class="md:hidden text-blue-600 transition-transform duration-300 ease-in-out"
                    :class="{ 'rotate-90': isSidebarOpen }"
                >
                    ☰
                </button>
            </div>
            <div class="bg-white p-4 shadow-md rounded-md">
                <form @submit.prevent="updateBlog" class="space-y-4">
                    <div>
                        <label class="block font-semibold mb-1">Judul</label>
                        <input v-model.trim="newBlog.title" type="text" placeholder="Masukkan judul blog..." class="w-full p-2 border rounded-md" required />
                    </div>
                    <div>
                        <label class="block font-semibold mb-1">Deskripsi</label>
                        <input v-model.trim="newBlog.description" type="text" placeholder="Masukkan deskripsi blog..." class="w-full p-2 border rounded-md" required />
                    </div>
                    <div>
                        <label class="block font-semibold mb-1">Kategori</label>
                        <select v-model="newBlog.category" class="w-full p-2 border rounded-md" required>
                            <option hidden value="">Pilih kategori...</option>
                            <option value="Nutrisi dan Diet">Nutrisi dan Diet</option>
                            <option value="Kesehatan Mental dan Mindfulness">Kesehatan Mental dan Mindfulness</option>
                            <option value="Kesehatan dan Kebugaran">Kesehatan dan Kebugaran</option>
                        </select>
                    </div>
                    <div>
                        <label class="block font-semibold mb-1">Isi</label>
                        <textarea v-model.trim="newBlog.content" placeholder="Masukkan isi blog..." class="w-full h-40 p-2 border rounded-md" required></textarea>
                    </div>
                    <div>
                        <label class="font-semibold w-full md:w-auto">Foto Header Blog</label>
                        <div 
                            @click="triggerFileInput" 
                            class="relative w-28 h-28 border-2 border-gray-300 flex items-center justify-center rounded-md cursor-pointer hover:bg-gray-100"
                        >
                            <input ref="fileHeaderInput" type="file" class="hidden" accept="image/*" @change="handleFileUpload" />
                            <img v-if="previewHeader" :src="previewHeader" class="absolute inset-0 w-full h-full object-cover rounded-md" />
                            <span v-else class="text-gray-400 text-3xl">+</span>
                        </div>
                    </div>
                    <div class="flex gap-4">
                        <router-link to="/admin-pg/blog" class="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600">
                            Batal
                        </router-link>
                        <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600">
                            Simpan
                        </button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
import axios from "axios";
import Sidebar from "../components/Sidebar.vue";

export default {
    components: { Sidebar },
    data() {
        return {
            isSidebarOpen: false,
            newBlog: {
                title: "",
                description: "",
                category: "",
                content: "",
                image: null,
            },
            previewHeader: null,
            id: this.$route.params.id,
        };
    },
    async mounted() {
        if (this.$route.state?.blog) {
            this.newBlog = { ...this.$route.state.blog };
            this.previewHeader = this.newBlog.image;
        } else {
            await this.fetchBlogById();
        }
    },
    methods: {
        async fetchBlogById() {
            try {
                const response = await axios.get(`http://localhost:8000/blog/${this.id}`);
                this.newBlog = response.data;
                this.previewHeader = this.newBlog.image 
                    ? `/uploads/${this.newBlog.image}` 
                    : "/default.jpg";
            } catch (error) {
                console.error("Gagal mengambil data blog:", error);
                alert("Gagal memuat data blog!");
            }
        },
        triggerFileInput() {
            this.$refs.fileHeaderInput.click();
        },
        handleFileUpload(event) {
            const file = event.target.files[0];
            if (file) {
                this.newBlog.image = file;
                const reader = new FileReader();
                reader.onload = (e) => {
                    this.previewHeader = e.target.result;
                };
                reader.readAsDataURL(file);
            }
        },
        async updateBlog() {
            if (!this.newBlog.title || !this.newBlog.description || !this.newBlog.category || !this.newBlog.content) {
                alert("Semua kolom wajib diisi!");
                return;
            }

            const formData = new FormData();
            formData.append("title", this.newBlog.title.trim());
            formData.append("description", this.newBlog.description.trim());
            formData.append("category", this.newBlog.category);
            formData.append("content", this.newBlog.content.trim());

            if (this.newBlog.image instanceof File) {
                formData.append("image", this.newBlog.image);
            }

            try {
                const response = await axios.patch(`http://localhost:8000/blog/update/${this.id}`, formData, {
                    headers: { "Content-Type": "multipart/form-data" },
                });

                if (response.status === 200) {
                    alert("Blog berhasil diperbarui!");
                    this.$router.push("/admin-pg/blog");
                } else {
                    throw new Error("Gagal memperbarui blog.");
                }
            } catch (error) {
                console.error("Error:", error.response?.data || error);
                alert("Gagal mengupdate blog! Cek koneksi API atau format data.");
            }
        },
    },
};
</script>
