// ducks: https://github.com/erikras/ducks-modular-redux

export const next = () => ({type: "EV_NEXT"});
export const prev = () => ({type: "EV_PREV"});
export const goto = (i) => ({type: "EV_GOTO", payload: i});

export default function reducer(state = {}, action) {
    const {index, history} = state;

    switch (action.type) {
        case "EV_NEXT":
            if (index === history.length - 1) {
                return state;
            }
            return {
                ...state,
                index: index + 1
            };
        case "EV_PREV":
            if (index === 0) {
                return state;
            }
            return {
                ...state,
                index: index - 1
            };
        case "EV_GOTO":
            return {
                ...state,
                index: action.payload
            };
    }
    return state;
}
