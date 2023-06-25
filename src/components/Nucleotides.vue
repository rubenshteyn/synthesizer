<script>
import {mapGetters} from "vuex";
import IconCogwheel from "@/components/icons/IconCogwheel.vue";

export default {
  name: "Nucleotides",
  components: {
    IconCogwheel
  },
  computed: {
    ...mapGetters({
      currentNucleotide: "synthesizer/getNucleotide"
    })
  }
}
</script>

<template>
  <div class="nucleotides">
    <div class="nucleotide">
      <img alt="Колба" src="@/assets/img/colba.png"/>
      <span :class="currentNucleotide === 'a' ? 'nucleotide__title active': 'nucleotide__title'">a</span>
      <IconCogwheel v-if="currentNucleotide === 'a'"/>
    </div>
    <div class="nucleotide">
      <img alt="Колба" src="@/assets/img/colba.png"/>
      <span :class="currentNucleotide === 't' ? 'nucleotide__title active': 'nucleotide__title'">t</span>
      <IconCogwheel v-if="currentNucleotide === 't'"/>
    </div>
    <div class="nucleotide">
      <img alt="Колба" src="@/assets/img/colba.png"/>
      <span :class="currentNucleotide === 'g' ? 'nucleotide__title active': 'nucleotide__title'">g</span>
      <IconCogwheel v-if="currentNucleotide === 'g'"/>
    </div>
    <div class="nucleotide">
      <img alt="Колба" src="@/assets/img/colba.png"/>
      <span :class="currentNucleotide === 'c' ? 'nucleotide__title active': 'nucleotide__title'">c</span>
      <IconCogwheel v-if="currentNucleotide === 'c'"/>
    </div>
  </div>
</template>

<style scoped>
.nucleotide .modal--cogwheel {
  position: absolute;
  fill: rgb(64, 199, 129);
  top: 23em;
}
</style>