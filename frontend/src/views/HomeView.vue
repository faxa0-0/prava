<script setup>
import StartButton from '@/components/StartQuizButton.vue';
import router from '@/router';

const quiz = localStorage.getItem('questions');
if (quiz) {
  router.push('/quiz');
}
</script>

<template>
  <div class="flex items-center justify-center min-h-screen bg-slate-100">
    <StartButton />
  </div>
</template>
