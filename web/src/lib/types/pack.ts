// Re-export generated types for convenience
export type { Pack, Calculation } from '$lib/api/generated/graphql';

// UI-specific types (not from GraphQL schema)
export interface PackInput {
	id: string;
	size: string;
	originalSize?: string;
	packUID?: string;
	saving: boolean;
	justSaved: boolean;
}
