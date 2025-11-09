import Dashboard from '../pages/DashboardPage.vue';
import userDashboard from '../pages/user/UserDashboard.vue';
import Sport from '../pages/user/sports/Sport.vue';
import Blog from '../pages/user/blogs/Blog.vue';
import Food from '../pages/user/foods/Food.vue';
import Recipe from '../pages/user/recipes/Recipe.vue';
import recipeDetail from '../pages/user/recipes/recipeDetail.vue';
import blogDetail from '../pages/user/blogs/blogDetail.vue';

const userRoutes = [
    { path: '/', name: 'Dashboard', component: Dashboard, meta: { title: 'Dashboard - IP-CORP' } },
    { path: '/dashboard', name: 'userDashboard', component: userDashboard, meta: { title: 'User Dashboard - IP-CORP' } },

    { path: '/blog', name: 'Blog', component: Blog, meta: { title: 'Blog - IP-CORP' } },
    { path: '/blog/detail', name: 'blogDetail', component: blogDetail, meta: { title: 'Blog-Detail - IP-CORP' } },

    { path: '/sport', name: 'Sport', component: Sport, meta: { title: 'Sport - IP-CORP' } },

    { path: '/food', name: 'Food', component: Food, meta: { title: 'Makanan - IP-CORP' } },
    
    { path: '/recipe', name: 'Recipe', component: Recipe, meta: { title: 'Recipe - IP-CORP' } },
    { path: '/recipe/detail', name: 'recipeDetail', component: recipeDetail, meta: { title: 'Recipe Detail - IP-CORP' } },
];

export default userRoutes;