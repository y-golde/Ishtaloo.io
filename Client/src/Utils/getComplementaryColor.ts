type Color = {
	r: number;
	g: number;
	b: number;
};

const getComplementary = (color: Color): Color => {
	const { r, g, b } = color;
	const maxMin = Math.max(r, b, g) + Math.min(r, b, g);
	const newR = maxMin - r;
	const newB = maxMin - b;
	const newG = maxMin - g;

	return { r: newR, b: newB, g: newG };
};

export default getComplementary;
