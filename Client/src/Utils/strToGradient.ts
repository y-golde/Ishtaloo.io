import getComplementary from './getComplementaryColor';

const squishNumber = (number: number) => {
	// we want the color not to be really bright nor really dark
	const minColorValue = 50;
	const maxColorValue = 200;

	/*
	 * if the number's rgb value is too large or too small -
	 * move it over to a normal value by adding or subtracting its diffrence from the allowed normal
	 * note that is does mean that the values 50-99 and 199-154 are more likely to appear than others (2 more times actually)
	 * I'll figure a way to address this inbalance in a next revision
	 */

	if (number > maxColorValue) {
		return maxColorValue - (255 - number);
	} else if (number < minColorValue) {
		return minColorValue + (50 - number);
	}
	return number;
};

const hashString = (s: string) =>
	Math.abs(s.split('').reduce((a, b) => ((a << 5) - a + b.charCodeAt(0)) | 0, 0)).toString(16);

const strToGradient = (str: string) => {
	const seed = hashString(str);

	// squish the color to a 'pretty' color - V.1
	const r = squishNumber(parseInt(seed.slice(0, 2), 16));
	const g = squishNumber(parseInt(seed.slice(2, 4), 16));
	const b = squishNumber(parseInt(seed.slice(4, 6), 16));
	const c = getComplementary({ r, g, b });

	return `linear-gradient(45deg, rgba(${r} , ${g}, ${b}), rgba(${c.r}, ${c.b}, ${c.g}))`;
};

export default strToGradient;
