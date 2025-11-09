export function isValidPackSize(size: string): boolean {
	const sizeStr = size?.toString().trim() || '';
	const sizeNum = parseInt(sizeStr, 10);
	return sizeStr !== '' && !isNaN(sizeNum) && sizeNum > 0;
}

export function parsePackSize(size: string): number | null {
	const sizeStr = size?.toString().trim() || '';
	const sizeNum = parseInt(sizeStr, 10);
	
	if (!sizeStr || isNaN(sizeNum) || sizeNum <= 0) {
		return null;
	}
	
	return sizeNum;
}
