import Dashboard from '../pages/DashboardPage.vue';
import Blog from '../pages/user/blogs/DashboardBlog.vue';
import Program from '../pages/user/programs/ProgramPage.vue';
import Sport from '../pages/user/sports/SportPage.vue';
import Food from '../pages/user/foods/FoodPage.vue';
import Recipe from '../pages/user/recipes/RecipePage.vue';
import BlogPage from '../pages/user/blogs/BlogPage.vue';

const userRoutes = [
    { path: '/', name: 'Dashboard', component: Dashboard, meta: { title: 'Dashboard - IP-CORP' } },

    { path: '/blog', name: 'Blog', component: Blog, meta: { title: 'Blog - IP-CORP' } },
    { path: '/blog-info', name: 'BlogPage', component: BlogPage, meta: { title: 'Blog-Info - IP-CORP' } },

    { path: '/sport', name: 'Sport', component: Sport, meta: { title: 'Sport - IP-CORP' } },

    { path: '/program', name: 'Program', component: Program, meta: { title: 'Program - IP-CORP' } },

    { path: '/food', name: 'Food', component: Food, meta: { title: 'Makanan - IP-CORP' } },
    
    { path: '/recipe', name: 'Recipe', component: Recipe, meta: { title: 'Recipe - IP-CORP' } },
];

export default userRoutes;