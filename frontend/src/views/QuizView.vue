<script setup>
import QuizNavigation from '@/components/QuizNavigation.vue';
</script>

<template>
  <div class="min-h-screen p-4 grid-cols-10 grid gap-4 text-white text-sm font-bold leading-6">
    <div
      class="rounded-lg shadow-lg row-span-2 col-span-8 text-gray-700 flex flex-col gap-6 py-6 px-4 overflow-y-auto max-h-[calc(100vh-2rem)] [&::-webkit-scrollbar]:hidden"
    >
      <img
        src="https://e-avtoshkola.uz/uploads/tests/questions/97/93-8.jpg"
        class="self-center h-2/4 w-2/4"
      />
      <div class="rounded-lg shadow-md py-4 px-3 text-xl">
        <p>Эта табличка означает:</p>
      </div>
      <div
        class="rounded-lg border-2 py-4 px-3 transition-colors duration-500 hover:bg-indigo-300 hover:text-white hover:border-indigo-400"
      >
        <p>Разрешает эвакуацию остановившихся транспортных средств в зоне действие знака</p>
      </div>
      <div
        class="rounded-lg border-2 py-4 px-3 transition-colors duration-500 hover:bg-indigo-300 hover:text-white hover:border-indigo-400"
      >
        <p>Запрещает эвакуацию остановившихся транспортных средств в зоне действие знака</p>
      </div>
      <div
        class="rounded-lg border-2 py-4 px-3 transition-colors duration-500 hover:bg-indigo-300 hover:text-white hover:border-indigo-400"
      >
        <p>Указывает на принудительную эвакуацию автомобилей, остановившихся в зоне действие знака</p>
      </div>
      <div class="hidden rounded-lg shadow-lg shadow-indigo-300 border-2 border-indigo-200 py-4 px-3 text-xl">
        <span>Объяснение</span>
        <p class="font-medium mt-2">
          A train passes a station platform in 36 seconds and a man standing on the platform in 20 seconds. If
          the speed of the train is 54 km/hr, what is the length of the platform? A train passes a station
          platform in 36 seconds and a man standing on the platform in 20 seconds. If the speed of the train
          is 54 km/hr, what is the length of the platform? A train passes a station platform in 36 seconds and
          a man standing on the platform in 20 seconds. If the speed of the train is 54 km/hr, what is the
          length of the platform? A train passes a station platform in 36 seconds and a man standing on the
          platform in 20 seconds. If the speed of the train is 54 km/hr, what is the length of the platform? A
          train passes a station platform in 36 seconds and a man standing on the platform in 20 seconds. If
          the speed of the train is 54 km/hr, what is the length of the platform? A train passes a station
          platform in 36 seconds and a man standing on the platform in 20 seconds. If the speed of the train
          is 54 km/hr, what is the length of the platform?
        </p>
      </div>
    </div>
    <QuizNavigation :answered-questions="10" :totalQuestions="20" />
  </div>
</template>
