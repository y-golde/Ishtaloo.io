const UPPER_BOUND = 200;
const LOWER_BOUND = 50;

const strToColor = (str: string) => {
	const hash = hashCode(str).toString(16);
	const r = squishColor(str.substring(0, 1));
	const g = squishColor(str.substring(2, 3));
	const b = squishColor(str.substring(4, 5));

	return `rgb(${r},${g},${b})`;
};

const squishColor = (color: string) => {
	const numeric = parseInt(color, 16);

	if (numeric > UPPER_BOUND) {
		return 256 - (numeric - UPPER_BOUND);
	} else if (numeric < LOWER_BOUND) {
		return 50 + (LOWER_BOUND - numeric);
	}
	return numeric;
};

export default strToColor;

const hashCode = (s: string) =>
	s.split('').reduce((a, b) => ((a << 5) - a + b.charCodeAt(0)) | 0, 0);
