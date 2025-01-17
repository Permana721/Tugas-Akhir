import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '../pages/DashboardPage.vue';
import Login from '../auth/Login.vue';
import Blog from '../pages/user/blogs/DashboardBlog.vue';
import Program from '../pages/user/programs/ProgramPage.vue';
import Sport from '../pages/user/sports/SportPage.vue';
import Food from '../pages/user/foods/FoodPage.vue';
import Recipe from '../pages/user/recipes/RecipePage.vue';
import BlogPage from '../pages/user/blogs/BlogPage.vue';
import Register from '../auth/Register.vue';


const routes = [
    // { path: '/', redirect: '/dashboard' }, 
    { path: '/', name: 'Dashboard', component: Dashboard, meta: { title: 'Dashboard - IP-CORP' } },
    { path: '/blog', name: 'Blog', component: Blog, meta: { title: 'Blog - IP-CORP' } },
    { path: '/blog-info', name: 'BlogPage', component: BlogPage, meta: { title: 'Blog-Info - IP-CORP' } },
    { path: '/sport', name: 'Sport', component: Sport, meta: { title: 'Sport - IP-CORP' } },
    { path: '/program', name: 'Program', component: Program, meta: { title: 'Program - IP-CORP' } },
    { path: '/food', name: 'Food', component: Food, meta: { title: 'Makanan - IP-CORP' } },
    { path: '/recipe', name: 'Recipe', component: Recipe, meta: { title: 'Recipe - IP-CORP' } },

    { path: '/login', name: 'Login', component: Login, meta: { title: 'Login' } },
    { path: '/register', name: 'Register', component: Register, meta: { title: 'Register' } },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title;
    } else {
        document.title = 'Default Title';
    }

    next();
});

export default router;