import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        isLoggedIn: false,
        token: null, // Tambahkan token
    }),
    actions: {
        login(token) {
            this.isLoggedIn = true;
            this.token = token; // Simpan token
        },
        logout() {
            this.isLoggedIn = false;
            this.token = null; // Hapus token
        },
        setLoginStatus(status) {
            this.isLoggedIn = status;
        },
    },
});