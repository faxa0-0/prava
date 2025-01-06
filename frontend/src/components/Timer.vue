<script setup>
import router from '@/router';
import { ref, onMounted, onUnmounted, watch } from 'vue';

const props = defineProps({
  initialMinutes: {
    type: Number,
    default: 20,
  },
  modelValue: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(['update:modelValue']);

const minutes = ref(props.initialMinutes);
const seconds = ref(0);
const isTimeUp = ref(props.modelValue);
let interval;

const saveTimerState = () => {
  localStorage.setItem(
    'timerState',
    JSON.stringify({
      minutes: minutes.value,
      seconds: seconds.value,
      lastSaved: Date.now(),
    })
  );
};

const loadTimerState = () => {
  const savedState = localStorage.getItem('timerState');
  if (savedState) {
    const { minutes: savedMinutes, seconds: savedSeconds, lastSaved } = JSON.parse(savedState);

    const elapsedSeconds = Math.floor((Date.now() - lastSaved) / 1000);
    const totalRemainingSeconds = savedMinutes * 60 + savedSeconds - elapsedSeconds;

    if (totalRemainingSeconds > 0) {
      minutes.value = Math.floor(totalRemainingSeconds / 60);
      seconds.value = totalRemainingSeconds % 60;
    } else {
      minutes.value = 0;
      seconds.value = 0;
      isTimeUp.value = true;
    }
  }
};

const startTimer = () => {
  interval = setInterval(() => {
    if (seconds.value === 0) {
      if (minutes.value > 0) {
        minutes.value--;
        seconds.value = 59;
      } else {
        isTimeUp.value = true;
        clearInterval(interval);
      }
    } else {
      seconds.value--;
    }
    saveTimerState();
  }, 1000);
};

const clearTimerState = () => {
  localStorage.removeItem('timerState');
};

watch(isTimeUp, (newVal) => {
  emit('update:modelValue', newVal);
  if (newVal) {
    clearTimerState();
    localStorage.removeItem('questions');
    router.push('/');
  }
});

onMounted(() => {
  loadTimerState();
  if (!isTimeUp.value) {
    startTimer();
  }
});

onUnmounted(() => {
  clearInterval(interval);
});
</script>

<template>
  <div class="flex flex-col items-center justify-center">
    <div class="text-2xl text-gray-700 font-bold">
      <span>{{ String(minutes).padStart(2, '0') }}</span
      >:
      <span>{{ String(seconds).padStart(2, '0') }}</span>
    </div>
  </div>
</template>
