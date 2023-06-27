export default {
    namespaced: true,
    state: {
        timer: 0,
        amountCompleted: 4,
        currentNucleotide: null,
        status: "Бездействует",
        synthesis: {
            synthesisHigh: [],
            synthesisMedium: [],
            synthesisLow: [],
        }
    },
    mutations: {
        ADD_STATUS(state, status) {
            state.status = status;
        },
        ADD_CURRENT_NUCLEOTIDE(state, currentNucleotide) {
            state.currentNucleotide = currentNucleotide;
        },
        ADD_SYNTHESIS_COMPLETED(state) {
            state.amountCompleted = state.amountCompleted + 1;
        },
        ADD_SYNTHESIS_TIMER(state) {
            const queues = Object.keys(state.synthesis)
            let arrNucleotides = [0]
            queues.forEach(queue => {
                state.synthesis[queue].forEach(synthes => {
                    arrNucleotides.push(synthes.nucleotides.length)
                })
            })
            state.timer = arrNucleotides.reduce((a, b) => a + b);
        },
        UPDATE_SYNTHESIS_TIMER(state) {
            if (state.timer === 0) {
                state.timer = 0
            }
            state.timer = state.timer - 1;
        },
        UPDATE_PRIORITY(state, data) {
            data.synthes.priority = data.futureQueue
            state.synthesis[data.futureQueue] = [...state.synthesis[data.futureQueue], data.synthes];
        },
        UPDATE_NUCLEOTIDES(state, data) {
            state.synthesis[data.queue] = state.synthesis[data.queue].filter((synthesis) => {
                if (synthesis.id === data.synthes.id) {
                    synthesis.nucleotides = data.nucleotides
                }
                return synthesis
            });
        },
        DELETE_SYNTHESIS(state, data) {
            state.synthesis[data.queue] = state.synthesis[data.queue].filter((synthesis) => synthesis.id !== data.synthes.id)
        },
        ADD_SYNTHESIS(state, data) {
            state.synthesis[data.synthes.priority] = [...state.synthesis[data.synthes.priority], data.synthes]
        },
        UPDATE_SYNTHESIS(state, data) {
            state.synthesis[data.queue] = state.synthesis[data.queue].filter((synthesis) => {
                if (synthesis.id === data.synthes.id) {
                    synthesis.work = data.synthes.work
                }
                return synthesis
            })
        },
    },
    actions: {
        addSynthesisMethod(context, data) {
            context.commit("ADD_SYNTHESIS", data)
        },

        updatePriorityMethod(context, data) {
            context.commit("DELETE_SYNTHESIS", data)
            context.commit("UPDATE_PRIORITY", data)
        },

        updateNucleotidesMethod(context, data) {
            context.commit("UPDATE_NUCLEOTIDES", data)
        },

        deleteSynthesisMethod(context, data) {
            context.commit("DELETE_SYNTHESIS", data)
        },

        updateStatusSynthesisMethod(context, data) {
            context.commit("UPDATE_SYNTHESIS", data)
        },

        subtractionTimeMethod(context) {
            context.commit("UPDATE_SYNTHESIS_TIMER")
        },

        updateTimerMethod(context) {
            context.commit("ADD_SYNTHESIS_TIMER")
        },

        updateSynthesisCompletedMethod(context) {
            context.commit("ADD_SYNTHESIS_COMPLETED")
        },

        updateCurrentNucleotideMethod(context, currentNucleotide) {
            context.commit("ADD_CURRENT_NUCLEOTIDE", currentNucleotide)
        },

        updateStatusMethod(context, status) {
            context.commit("ADD_STATUS", status)
        },
    },

    getters: {
        getSynthesis(state) {
            return state.synthesis
        },
        getTimer(state) {
            return state.timer
        },
        getSynthesisCompleted(state) {
            return state.amountCompleted
        },
        getNucleotide(state) {
            return state.currentNucleotide
        },
        getStatus(state) {
            return state.status
        },
    }
};