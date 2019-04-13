<template>
  <div class="spinner-wrapper">
    <div class="spinner">
      <svg width="100%" height="100%" viewBox="0 0 66 66" xmlns="http://www.w3.org/2000/svg" fit="" preserveAspectRatio="xMidYMid meet" focusable="false">
        <circle class="path" fill="none" stroke-width="6" stroke-linecap="round" cx="33" cy="33" r="30"></circle>
      </svg>
    </div>

    <h2 class="spinner-text" v-if="text">{{ text }}</h2>
  </div>
</template>

<script>
export default {
  name: 'Spinner',
  props: ['text']
}
</script>

<style lang="scss">
@import '@/assets/styles/bulma/_utilities.scss';

.spinner-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}

.spinner {
  display: inline-block;
  width: 26px;
  height: 26px;
  margin-bottom: 20px;

  &:last-child {
    margin-bottom: 0;
  }

  > svg {
    animation: 1.4s linear infinite rotator;

    .path {
      stroke-dasharray: 187;
      stroke-dashoffset: 0;
      transform-origin: center;
      stroke: $corporate-green;
      animation: 1.4s ease-in-out infinite dash;
    }
  }
}

@keyframes rotator {
  0% {
    transform: rotate(0);
  }

  100% {
    transform: rotate(270deg);
  }
}

@keyframes dash {
  0% {
    stroke-dashoffset: 187
  }

  50% {
    stroke-dashoffset: 46.75;
    transform: rotate(135deg)
  }

  100% {
    stroke-dashoffset: 187;
    transform: rotate(450deg)
  }
}

.spinner-text {
  font-size: $size-normal;
  color: $grey-dark;
}
</style>
