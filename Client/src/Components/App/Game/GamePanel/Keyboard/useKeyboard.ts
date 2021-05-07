const UseKeyboard = (disabledChars: string[]) => {
	const isDisabled = (char: string) => {
		return disabledChars.indexOf(char) !== -1;
	};

	return {
		isDisabled,
	};
};

export default UseKeyboard;
