<script>
import {mapGetters} from "vuex";
import IconCogwheel from "@/components/icons/IconCogwheel.vue";

export default {
  name: "Nucleotides",
  data() {
    return {
      nucleotides: ['a', 't', 'g', 'c']
    }
  },
  components: {
    IconCogwheel
  },
  computed: {
    ...mapGetters({
      currentNucleotide: "synthesizer/getNucleotide",
    })
  }
}
</script>

<template>
  <div class="nucleotides">
    <div v-for="nucleotide in nucleotides" class="nucleotide">
      <img alt="Колба" src="@/assets/img/colba.png"/>
      <span
          :class="currentNucleotide === nucleotide ? 'nucleotide__title active': 'nucleotide__title'">
          {{ nucleotide }}
      </span>
      <IconCogwheel v-if="currentNucleotide === nucleotide"/>
    </div>
  </div>
</template>

<style scoped>
.nucleotide .modal--cogwheel {
  position: absolute;
  fill: rgb(64, 199, 129);
  top: 19em;
}
</style>