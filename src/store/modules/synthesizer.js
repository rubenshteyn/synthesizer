export default {
    namespaced: true,
    state: {
        timer: 0,
        amountCompleted: 0,
        currentNucleotide: null,
        status: "Бездействует",
        synthesis: {
            synthesisHigh: [],
            synthesisMedium: [],
            synthesisLow: [],
        },
    },
    mutations: {
        ADD_STATUS(state, status) {
            state.status = status;
        },
        ADD_CURRENT_NUCLEOTIDE(state, currentNucleotide) {
            state.currentNucleotide = currentNucleotide;
        },
        ADD_SYNTHESIS_COMPLETED(state, amountCompleted) {
            state.amountCompleted = amountCompleted;
        },
        ADD_SYNTHESIS_TIMER(state, timer) {
            state.timer = timer;
        },
        ADD_SYNTHESIS_HIGH(state, synthesis) {
            state.synthesis.synthesisHigh = synthesis;
        },
        ADD_SYNTHESIS_MEDIUM(state, synthesis) {
            state.synthesis.synthesisMedium = synthesis;
        },
        ADD_SYNTHESIS_LOW(state, synthesis) {
            state.synthesis.synthesisLow = synthesis;
        }
    },
    actions: {
        addSynthesisMethod(context, form) {
            let updateState
            if (form.priority === "high") {
                updateState = [...context.state.synthesis.synthesisHigh, form]
                return context.commit("ADD_SYNTHESIS_HIGH", updateState)
            }
            if (form.priority === "medium") {
                updateState = [...context.state.synthesis.synthesisMedium, form]
                return context.commit("ADD_SYNTHESIS_MEDIUM", updateState)
            }
            if (form.priority === "low") {
                updateState = [...context.state.synthesis.synthesisLow, form]
                return context.commit("ADD_SYNTHESIS_LOW", updateState)
            }
        },

        updatePriorityMethod(context, data) {
            if (data.synthes.priority === "high" && data.futurePriority === "low") {
                data.synthes.priority = data.futurePriority
                context.commit("ADD_SYNTHESIS_LOW", [...context.state.synthesis.synthesisLow, data.synthes])
                context.commit("ADD_SYNTHESIS_HIGH", context.state.synthesis.synthesisHigh.filter((synthesis) => synthesis.id !== data.synthes.id))
            }

            if (data.synthes.priority === "high" && data.futurePriority === "medium") {
                data.synthes.priority = data.futurePriority
                context.commit("ADD_SYNTHESIS_MEDIUM", [...context.state.synthesis.synthesisMedium, data.synthes])
                context.commit("ADD_SYNTHESIS_HIGH", context.state.synthesis.synthesisHigh.filter((synthesis) => synthesis.id !== data.synthes.id))
            }

            if (data.synthes.priority === "medium" && data.futurePriority === "high") {
                data.synthes.priority = data.futurePriority
                context.commit("ADD_SYNTHESIS_HIGH", [...context.state.synthesis.synthesisHigh, data.synthes])
                context.commit("ADD_SYNTHESIS_MEDIUM", context.state.synthesis.synthesisMedium.filter((synthesis) => synthesis.id !== data.synthes.id))
            }

            if (data.synthes.priority === "medium" && data.futurePriority === "low") {
                data.synthes.priority = data.futurePriority
                context.commit("ADD_SYNTHESIS_LOW", [...context.state.synthesis.synthesisLow, data.synthes])
                context.commit("ADD_SYNTHESIS_MEDIUM", context.state.synthesis.synthesisMedium.filter((synthesis) => synthesis.id !== data.synthes.id))
            }

            if (data.synthes.priority === "low" && data.futurePriority === "medium") {
                data.synthes.priority = data.futurePriority
                context.commit("ADD_SYNTHESIS_MEDIUM", [...context.state.synthesis.synthesisMedium, data.synthes])
                context.commit("ADD_SYNTHESIS_LOW", context.state.synthesis.synthesisLow.filter((synthesis) => synthesis.id !== data.synthes.id))
            }

            if (data.synthes.priority === "low" && data.futurePriority === "high") {
                data.synthes.priority = data.futurePriority
                context.commit("ADD_SYNTHESIS_HIGH", [...context.state.synthesis.synthesisHigh, data.synthes])
                context.commit("ADD_SYNTHESIS_LOW", context.state.synthesis.synthesisLow.filter((synthesis) => synthesis.id !== data.synthes.id))
            }
        },

        updateNucleotidesMethod(context, data) {
            if (data.synthes.priority === "high") {
                context.commit("ADD_SYNTHESIS_HIGH", context.state.synthesis.synthesisHigh.filter((synthesis) => {
                    if (synthesis.id === data.synthes.id) {
                        synthesis.nucleotides = data.nucleotides
                    }
                    return synthesis
                }))
            }
            if (data.synthes.priority === "medium") {
                context.commit("ADD_SYNTHESIS_MEDIUM", context.state.synthesis.synthesisMedium.filter((synthesis) => {
                    if (synthesis.id === data.synthes.id) {
                        synthesis.nucleotides = data.nucleotides
                    }
                    return synthesis
                }))
            }
            if (data.synthes.priority === "low") {
                context.commit("ADD_SYNTHESIS_LOW", context.state.synthesis.synthesisLow.filter((synthesis) => {
                    if (synthesis.id === data.synthes.id) {
                        synthesis.nucleotides = data.nucleotides
                    }
                    return synthesis
                }))
            }
        },

        deleteSynthesisMethod(context, form) {
            let updateState
            if (form.priority === "high") {
                updateState = context.state.synthesis.synthesisHigh.filter((synthesis) => synthesis.id !== form.id)
                return context.commit("ADD_SYNTHESIS_HIGH", updateState)
            }
            if (form.priority === "medium") {
                updateState = context.state.synthesis.synthesisMedium.filter((synthesis) => synthesis.id !== form.id)
                return context.commit("ADD_SYNTHESIS_MEDIUM", updateState)
            }
            if (form.priority === "low") {
                updateState = context.state.synthesis.synthesisLow.filter((synthesis) => synthesis.id !== form.id)
                return context.commit("ADD_SYNTHESIS_LOW", updateState)
            }
        },

        updateStatusSynthesisMethod(context, data) {
            if (data.synthes.priority === "high") {
                context.commit("ADD_SYNTHESIS_HIGH", context.state.synthesis.synthesisHigh.filter((synthesis) => {
                    if (synthesis.id === data.synthes.id) {
                        synthesis.work = data.work
                    }
                    return synthesis
                }))
            }
            if (data.synthes.priority === "medium") {
                context.commit("ADD_SYNTHESIS_MEDIUM", context.state.synthesis.synthesisMedium.filter((synthesis) => {
                    if (synthesis.id === data.synthes.id) {
                        synthesis.work = data.work
                    }
                    return synthesis
                }))
            }
            if (data.synthes.priority === "low") {
                context.commit("ADD_SYNTHESIS_LOW", context.state.synthesis.synthesisLow.filter((synthesis) => {
                    if (synthesis.id === data.synthes.id) {
                        synthesis.work = data.work
                    }
                    return synthesis
                }))
            }
        },

        subtractionTimeMethod(context) {
            let stateTimer = context.state.timer
            if (stateTimer === 0) {
                return context.commit("ADD_SYNTHESIS_TIMER", 0)
            }
            return context.commit("ADD_SYNTHESIS_TIMER", stateTimer - 1)
        },

        updateTimerMethod(context) {
            let arrNucleotides = [0]
            context.state.synthesis.synthesisHigh.forEach(synthes => {
                arrNucleotides.push(synthes.nucleotides.length)
            })
            context.state.synthesis.synthesisMedium.forEach(synthes => {
                arrNucleotides.push(synthes.nucleotides.length)
            })
            context.state.synthesis.synthesisLow.forEach(synthes => {
                arrNucleotides.push(synthes.nucleotides.length)
            })
            return context.commit("ADD_SYNTHESIS_TIMER", arrNucleotides.reduce((a, b) => a + b))
        },

        updateSynthesisCompletedMethod(context) {
            return context.commit("ADD_SYNTHESIS_COMPLETED", context.state.amountCompleted + 1)
        },

        updateCurrentNucleotideMethod(context, currentNucleotide) {
            return context.commit("ADD_CURRENT_NUCLEOTIDE", currentNucleotide)
        },

        updateStatusMethod(context, status) {
            return context.commit("ADD_STATUS", status)
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