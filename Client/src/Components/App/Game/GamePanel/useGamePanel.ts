import BLANK_SYMBOLS from './blankSymbols';

const UseGamePanel = () => {
    const isBlankSymbol = (char : string) => {
        return BLANK_SYMBOLS.indexOf(char) !== -1   
    }

    const getCorrectCharsByWord = (word: string) => {
        return word.split(' ').filter(char => !isBlankSymbol(char));
    }

    return {
        getCorrectCharsByWord
    }
}

export default UseGamePanel;