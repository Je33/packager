export const SUCCESS_FEEDBACK_DURATION = 2000; // 2 seconds

export const ERROR_MESSAGES = {
	INVALID_PACK_SIZE: 'Please enter a valid pack size',
	INVALID_ITEMS: 'Please enter a valid number of items',
	FETCH_PACKS_FAILED: 'Failed to fetch packs',
	CREATE_PACK_FAILED: 'Failed to create pack',
	UPDATE_PACK_FAILED: 'Failed to update pack',
	DELETE_PACK_FAILED: 'Failed to delete pack',
	CALCULATE_FAILED: 'Failed to calculate packs'
} as const;
