import BLANK_SYMBOLS from './blankSymbols';

const UseGamePanel = () => {
	const isBlankSymbol = (char: string) => {
		return BLANK_SYMBOLS.indexOf(char) !== -1;
	};

	const getCorrectCharsByWord = (word: string) => {
		return Boolean(word)
			? word.split(' ').filter((char) => !isBlankSymbol(char))
			: [];
	};

	const fetchGameInfo = async () => {
		// this is a stub that will be changed when SSE will be implemented 
		return {
			word: '_ _ P _ - _ _ _ A _',
			wrongGuesses: ['Q', 'M', 'Z']
		}
	}

	const getGameInfo = async () => {
		const gameInfo = await fetchGameInfo();
		return gameInfo;
	}

	return {
		getCorrectCharsByWord,
		getGameInfo
	};
};

export default UseGamePanel;
