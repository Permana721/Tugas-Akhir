<template>
    <Navbar />
    <div class="min-h-screen pt-28 py-6">
        <div class="max-w-4xl mx-auto bg-white rounded-lg p-6">
            <div>
                <h1 class="text-3xl font-bold text-gray-800 mb-4">
                    Augie Nieto: 2024 Don Wildman Commitment to Excellence Award Winner
                </h1>
                <p class="text-sm text-gray-600">
                    Admin | <a href="#" class="text-blue-600 underline">Tetap sehat dengan IP-CORP</a>
                </p>
            </div>
            <div class="mt-6">
                <img
                    src="/img/background/olahraga.jpg"
                    alt="Augie Nieto"
                    class="rounded-lg object-cover"
                />
            </div>
            <div class="mt-6 space-y-4 text-gray-700 text-lg">
                <p>
                    In the world of fitness, few names resonate as profoundly as Don Wildman’s. A trailblazing visionary who
                    revolutionized the industry, Wildman reimagined the concept of health clubs, making fitness accessible and
                    mainstream in many inspiring ways.
                </p>
                <p>
                    Known for his legendary high-intensity workouts and as an Ironman age group winner, Don is credited with
                    introducing music to group fitness classes, a pivotal change that significantly enhanced the welcoming
                    atmosphere of gyms, particularly for women.
                </p>
            </div>
            <div class="mt-8">
                <img
                    src="/img/background/olahraga.jpg"
                    alt="Tribute to Augie Nieto"
                    class="rounded-lg object-cover"
                />
            </div>
            <div class="mt-10 border-t border-gray-200 pt-6">
                <h3 class="text-xl font-semibold text-gray-800 mb-4">Related Posts</h3>
                <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
                    <RouterLink
                    to="/blog/discussing-alcohol-nutrition"
                    class="bg-white rounded-lg shadow-md overflow-hidden"
                    >
                    <img
                        src="/img/background/olahraga.jpg"
                        alt="Discussing Alcohol & Nutrition"
                        class="w-full h-28 object-cover"
                    />
                    <div class="p-2">
                        <h4 class="text-sm font-bold text-gray-800">
                        Discussing Alcohol & Nutrition: Making Mindful Choices
                        </h4>
                    </div>
                    </RouterLink>
                    <RouterLink
                    to="/blog/calculate-rmr"
                    class="bg-white rounded-lg shadow-md overflow-hidden"
                    >
                    <img
                        src="/img/background/olahraga.jpg"
                        alt="Here's How to Calculate & Improve Your RMR"
                        class="w-full h-28 object-cover"
                    />
                    <div class="p-2">
                        <h4 class="text-sm font-bold text-gray-800">
                        Here's How to Calculate & Improve Your RMR
                        </h4>
                    </div>
                    </RouterLink>
                    <RouterLink
                    to="/blog/healthy-meal-prep"
                    class="bg-white rounded-lg shadow-md overflow-hidden"
                    >
                    <img
                        src="/img/background/olahraga.jpg"
                        alt="4 Easy Steps to Meal Prep"
                        class="w-full h-28 object-cover"
                    />
                    <div class="p-2">
                        <h4 class="text-sm font-bold text-gray-800">
                        A Quick & Healthy Guide: 4 Easy Steps to Meal Prep
                        </h4>
                    </div>
                    </RouterLink>
                </div>
            </div>
        </div>
    </div>
    <Footer />
</template>

<script>
import Navbar from '../components/Navbar.vue';
import Footer from '../components/Footer.vue';

export default {
    components: {
        Navbar,
        Footer,
    },
};
</script>