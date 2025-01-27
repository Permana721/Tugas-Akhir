import Login from '../auth/Login.vue';
import Register from '../auth/Register.vue';

const authRoutes = [
    { path: '/login', name: 'Login', component: Login, meta: { title: 'Login' } },
    { path: '/register', name: 'Register', component: Register, meta: { title: 'Register' } },
]

export default authRoutes;