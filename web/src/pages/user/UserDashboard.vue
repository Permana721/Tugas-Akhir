<template>
  <Navbar />
  <div class="min-h-screen pt-28 bg-gray-100 p-6">
    <div class="max-w-5xl mx-auto bg-white rounded-lg shadow-lg p-6">
      <div class="border-b pb-4 mb-6 flex items-center justify-between">
        <div>
          <h1 class="text-xl font-semibold text-gray-800 flex items-center">
            <span>Makanan Harian Saya</span>
            <input
              type="date"
              v-model="selectedDate"
              class="ml-2 border rounded-lg px-2 py-1 text-sm text-gray-600 focus:ring-2 focus:ring-blue-500"
            />
          </h1>
        </div>
      </div>

      <div class="space-y-6">
        <div
          v-for="section in [
          { id: 1, title: 'Makan Pagi' },
          { id: 2, title: 'Makan Siang' },
          { id: 3, title: 'Makan Sore' },
          { id: 4, title: 'Cemilan' }
          ]"
          :key="section.id"
        >
          <h2 class="text-lg font-semibold text-gray-800 mb-2">{{ section.title }}</h2>
          <button class="w-full py-2 px-4 text-blue-600 rounded-lg border-2 text-left font-bold border-blue-600 hover:bg-gray-100" @click="addItem(section.title)">
            + Tambah Item
          </button>
        </div>
      </div>

      <div class="border-b pb-4 mb-6 mt-7 flex items-center justify-between">
        <div>
          <h1 class="text-xl font-semibold text-gray-800 flex items-center">
            <span>Catatan Harian Saya</span>
          </h1>
        </div>
      </div>

      <div class="space-y-6 mb-6">
        <div
          v-for="section in [
          { id: 1, title: 'Kegiatan Latihan' },
          ]"
          :key="section.id"
        >
          <h2 class="text-lg font-semibold text-gray-800 mb-2">{{ section.title }}</h2>
          <button class="w-full py-2 px-4 text-blue-600 rounded-lg border-2 text-left font-bold border-blue-600 hover:bg-gray-100" @click="addItem(section.title)">
            + Tambah Latihan
          </button>
        </div>
      </div>

      <div>
        <h2 class="text-lg font-semibold mt-4 text-gray-800 mb-4">Konsumsi Air</h2>
        <div class="bg-blue-50 p-4 rounded-lg">
          <div class="flex flex-col sm:flex-row justify-between items-center sm:items-start mb-4">
            <div class="flex-1 mb-4 sm:mb-0">
              <div class="flex justify-between text-sm text-gray-600 mb-1">
                <span>Terkonsumsi: {{ waterConsumed }} ml</span>
                <span class="mr-5">Target: {{ waterGoal }} ml</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2.5 relative overflow-hidden">
                <div 
                  class="bg-blue-500 h-2.5 rounded-full absolute top-0 left-0 transition-all duration-300"
                  :style="{ width: `${Math.min((waterConsumed / waterGoal) * 100, 100)}%` }"
                ></div>
                <div 
                  v-if="waterConsumed > waterGoal"
                  class="bg-red-500 h-2.5 rounded-full absolute top-0 transition-all duration-300"
                  :style="{
                    left: `${100}%`,
                    width: `${Math.min(((waterConsumed - waterGoal) / waterGoal) * 100, 100)}%`
                  }"
                ></div>
              </div>
            </div>
            <button
              @click="resetWater"
              class="px-3 py-1 text-sm bg-red-100 text-red-600 rounded-lg hover:bg-red-200"
            >
              Reset
            </button>
          </div>

          <div 
            class="mb-4 p-2 rounded-lg text-sm text-center"
            :class="{
              'bg-green-100 text-green-700': waterConsumed >= waterGoal && waterConsumed <= waterGoal * 1.5,
              'bg-red-100 text-red-700': waterConsumed > waterGoal * 1.5,
              'bg-blue-100 text-blue-700': waterConsumed < waterGoal
            }"
          >
            <template v-if="waterConsumed < waterGoal">
              🎯 Target hari ini: {{ waterGoal }} ml ({{ waterGoal - waterConsumed }} ml tersisa)
            </template>
            <template v-else-if="waterConsumed <= waterGoal * 1.5">
              🎉 Target tercapai! (+{{ waterConsumed - waterGoal }} ml)
            </template>
            <template v-else>
              ⚠️ Konsumsi melebihi {{ Math.round((waterConsumed / waterGoal) * 100) }}% target!
            </template>
          </div>

          <div class="flex flex-col sm:flex-row gap-2 mb-4">
            <button
              v-for="amount in [250, 500, 1000]"
              :key="amount"
              @click="addWater(amount)"
              class="flex-1 px-4 py-2 bg-white text-blue-600 rounded-lg border border-blue-200 hover:bg-blue-100 text-sm sm:text-base"
            >
              +{{ amount }}ml
            </button>
          </div>

          <div class="flex flex-col sm:flex-row gap-2">
            <input
              type="number"
              v-model.number="customWater"
              placeholder="Jumlah custom (ml)"
              class="flex-1 px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500 text-sm sm:text-base"
              min="1"
            >
            <button
              @click="addWater(customWater)"
              class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 text-sm sm:text-base"
            >
              Tambah
            </button>
          </div>
        </div>
      </div>

      <div class="mt-8 bg-gray-50 p-4 rounded-lg shadow">
        <h2 class="text-lg font-semibold text-gray-800 mb-2">Ringkasan Hari Ini:</h2>
        <p class="text-sm text-gray-600 mb-2">Berikut ringkasan apa yang Anda makan hari ini:</p>
        <div class="flex flex-col lg:flex-row justify-between items-center space-y-6 lg:space-y-0 lg:space-x-8">
          <div class="flex-1 w-full lg:w-auto">
            <div class="flex justify-between text-gray-700 mb-2">
              <span>Total Kalori</span>
              <span>{{ 2000 }} kkal</span>
            </div>
            <div class="flex justify-between text-gray-700 mb-2">
              <div class="flex items-center">
                <div class="w-4 h-4 bg-green-500 mr-2"></div> 
                <span>Karbohidrat</span>
              </div>
              <span>{{ 50 }}%</span>
            </div>
            <div class="flex justify-between text-gray-700 mb-2">
              <div class="flex items-center">
                <div class="w-4 h-4 bg-orange-500 mr-2"></div> 
                <span>Lemak</span>
              </div>
              <span>{{ 30 }}%</span>
            </div>
            <div class="flex justify-between text-gray-700 mb-2">
              <div class="flex items-center">
                <div class="w-4 h-4 bg-blue-500 mr-2"></div> 
                <span>Protein</span>
              </div>
              <span>{{ 20 }}%</span>
            </div>
            <div class="flex justify-between text-gray-700 mb-2">
              <div class="flex items-center">
                <div class="w-4 h-4 bg-red-500 mr-2"></div> 
                <span>Gula</span>
              </div>
              <span>{{ 10 }}%</span>
            </div>
            <div class="flex justify-between text-gray-700 mb-2">
              <div class="flex items-center">
                <div class="w-4 h-4 bg-purple-500 mr-2"></div> 
                <span>Sodium</span>
              </div>
              <span>{{ 5 }}%</span>
            </div>
          </div>

          <div class="flex-1 w-full lg:w-auto">
            <canvas id="summaryChart" class="w-full h-48 sm:h-36"></canvas>
          </div>
        </div>
      </div>
    </div>
  </div>
  <Footer />
</template>

<script>
import Navbar from '../../components/Navbar.vue';
import Footer from '../../components/Footer.vue';
import { Chart as ChartJS, Title, Tooltip, Legend, ArcElement, CategoryScale, LinearScale, PieController } from 'chart.js';

export default {
  data() {
    return {
      waterConsumed: 0,
      waterGoal: 2000,
      customWater: null,
      selectedDate: new Date().toISOString().substring(0, 10),
      chart: null,
    };
  },
  methods: {
    addWater(amount) {
      if (amount && amount > 0) {
        this.waterConsumed += amount;
        this.customWater = null;
      }
    },
    resetWater() {
      this.waterConsumed = 0;
    },
    createChart() {
        ChartJS.register(ArcElement, Tooltip, Legend, Title, CategoryScale, LinearScale, PieController);

        const ctx = document.getElementById('summaryChart').getContext('2d');
        this.chart = new ChartJS(ctx, {
          type: 'pie',
          data: {
          labels: ['Karbohidrat', 'Lemak', 'Protein', 'Gula', 'Sodium'],
          datasets: [
              {
              label: 'Ringkasan Nutrisi',
              data: [50, 30, 20, 10, 5],
              backgroundColor: ['#4caf50', '#ff9800', '#2196f3', '#e91e63', '#9c27b0'],
              borderWidth: 1,
              },
          ],
          },
          options: {
          responsive: true,
          maintainAspectRatio: false, 
          plugins: {
              legend: {
              display: false, 
              },
            },
          },
        });
    },
  },
  mounted() {
    this.createChart();
  },
  components: {
    Navbar,
    Footer
  },
};
</script>
