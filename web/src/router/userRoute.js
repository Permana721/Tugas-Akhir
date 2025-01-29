import Dashboard from '../pages/DashboardPage.vue';
import UserDashboard from '../pages/user/UserDashboard.vue';
import Blog from '../pages/user/blogs/DashboardBlog.vue';
import Sport from '../pages/user/sports/SportPage.vue';
import Food from '../pages/user/foods/FoodPage.vue';
import Recipe from '../pages/user/recipes/RecipePage.vue';
import BlogPage from '../pages/user/blogs/BlogPage.vue';

const userRoutes = [
    { path: '/', name: 'Dashboard', component: Dashboard, meta: { title: 'Dashboard - IP-CORP' } },
    { path: '/dashboard', name: 'UserDashboard', component: UserDashboard, meta: { title: 'User Dashboard - IP-CORP' } },

    { path: '/blog', name: 'Blog', component: Blog, meta: { title: 'Blog - IP-CORP' } },
    { path: '/blog-info', name: 'BlogPage', component: BlogPage, meta: { title: 'Blog-Info - IP-CORP' } },

    { path: '/sport', name: 'Sport', component: Sport, meta: { title: 'Sport - IP-CORP' } },

    { path: '/food', name: 'Food', component: Food, meta: { title: 'Makanan - IP-CORP' } },
    
    { path: '/recipe', name: 'Recipe', component: Recipe, meta: { title: 'Recipe - IP-CORP' } },
];

export default userRoutes;