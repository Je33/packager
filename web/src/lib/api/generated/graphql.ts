/* eslint-disable */
import type { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = T | null | undefined;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = {
	[_ in K]?: never;
};
export type Incremental<T> =
	| T
	| { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
	ID: { input: string; output: string };
	String: { input: string; output: string };
	Boolean: { input: boolean; output: boolean };
	Int: { input: number; output: number };
	Float: { input: number; output: number };
};

/**  Calculation of packages and items  */
export type Calculation = {
	__typename: 'Calculation';
	/**  Amount of items to fill  */
	Items: Scalars['Int']['output'];
	/**  Package capacity  */
	PackSize: Scalars['Int']['output'];
	/**  Package identifier  */
	PackUID: Scalars['String']['output'];
	/**  Amount of packages  */
	Quantity: Scalars['Int']['output'];
};

export type Mutation = {
	__typename: 'Mutation';
	packCreate: PackCreateResponse;
	packDelete: PackDeleteResponse;
	packUpdate: PackUpdateResponse;
};

export type MutationPackCreateArgs = {
	input: PackCreateRequest;
};

export type MutationPackDeleteArgs = {
	input: PackDeleteRequest;
};

export type MutationPackUpdateArgs = {
	input: PackUpdateRequest;
};

/**  Package type  */
export type Pack = {
	__typename: 'Pack';
	/**  Package size  */
	Size: Scalars['Int']['output'];
	/**  Package identifier  */
	UID: Scalars['String']['output'];
};

export type PackCalculateRequest = {
	items: Scalars['Int']['input'];
};

export type PackCalculateResponse = {
	__typename: 'PackCalculateResponse';
	calculations: Array<Maybe<Calculation>>;
};

export type PackCreateRequest = {
	size: Scalars['Int']['input'];
};

export type PackCreateResponse = {
	__typename: 'PackCreateResponse';
	pack: Pack;
};

export type PackDeleteRequest = {
	uid: Scalars['String']['input'];
};

export type PackDeleteResponse = {
	__typename: 'PackDeleteResponse';
	pack: Pack;
};

export type PackGetAllRequest = {
	limit?: InputMaybe<Scalars['Int']['input']>;
	page?: InputMaybe<Scalars['Int']['input']>;
};

export type PackGetAllResponse = {
	__typename: 'PackGetAllResponse';
	packs: Array<Maybe<Pack>>;
};

export type PackUpdateRequest = {
	size: Scalars['Int']['input'];
	uid: Scalars['String']['input'];
};

export type PackUpdateResponse = {
	__typename: 'PackUpdateResponse';
	pack: Pack;
};

export type Query = {
	__typename: 'Query';
	packCalculate: PackCalculateResponse;
	packGetAll: PackGetAllResponse;
};

export type QueryPackCalculateArgs = {
	input: PackCalculateRequest;
};

export type QueryPackGetAllArgs = {
	input: PackGetAllRequest;
};

export type GetPacksQueryVariables = Exact<{
	input: PackGetAllRequest;
}>;

export type GetPacksQuery = {
	__typename: 'Query';
	packGetAll: {
		__typename: 'PackGetAllResponse';
		packs: Array<{ __typename: 'Pack'; UID: string; Size: number } | null>;
	};
};

export type CreatePackMutationVariables = Exact<{
	input: PackCreateRequest;
}>;

export type CreatePackMutation = {
	__typename: 'Mutation';
	packCreate: {
		__typename: 'PackCreateResponse';
		pack: { __typename: 'Pack'; UID: string; Size: number };
	};
};

export type UpdatePackMutationVariables = Exact<{
	input: PackUpdateRequest;
}>;

export type UpdatePackMutation = {
	__typename: 'Mutation';
	packUpdate: {
		__typename: 'PackUpdateResponse';
		pack: { __typename: 'Pack'; UID: string; Size: number };
	};
};

export type DeletePackMutationVariables = Exact<{
	input: PackDeleteRequest;
}>;

export type DeletePackMutation = {
	__typename: 'Mutation';
	packDelete: {
		__typename: 'PackDeleteResponse';
		pack: { __typename: 'Pack'; UID: string; Size: number };
	};
};

export type CalculatePacksQueryVariables = Exact<{
	input: PackCalculateRequest;
}>;

export type CalculatePacksQuery = {
	__typename: 'Query';
	packCalculate: {
		__typename: 'PackCalculateResponse';
		calculations: Array<{
			__typename: 'Calculation';
			PackUID: string;
			PackSize: number;
			Quantity: number;
		} | null>;
	};
};

export const GetPacksDocument = {
	kind: 'Document',
	definitions: [
		{
			kind: 'OperationDefinition',
			operation: 'query',
			name: { kind: 'Name', value: 'GetPacks' },
			variableDefinitions: [
				{
					kind: 'VariableDefinition',
					variable: { kind: 'Variable', name: { kind: 'Name', value: 'input' } },
					type: {
						kind: 'NonNullType',
						type: { kind: 'NamedType', name: { kind: 'Name', value: 'PackGetAllRequest' } }
					}
				}
			],
			selectionSet: {
				kind: 'SelectionSet',
				selections: [
					{
						kind: 'Field',
						name: { kind: 'Name', value: 'packGetAll' },
						arguments: [
							{
								kind: 'Argument',
								name: { kind: 'Name', value: 'input' },
								value: { kind: 'Variable', name: { kind: 'Name', value: 'input' } }
							}
						],
						selectionSet: {
							kind: 'SelectionSet',
							selections: [
								{
									kind: 'Field',
									name: { kind: 'Name', value: 'packs' },
									selectionSet: {
										kind: 'SelectionSet',
										selections: [
											{ kind: 'Field', name: { kind: 'Name', value: 'UID' } },
											{ kind: 'Field', name: { kind: 'Name', value: 'Size' } }
										]
									}
								}
							]
						}
					}
				]
			}
		}
	]
} as unknown as DocumentNode<GetPacksQuery, GetPacksQueryVariables>;
export const CreatePackDocument = {
	kind: 'Document',
	definitions: [
		{
			kind: 'OperationDefinition',
			operation: 'mutation',
			name: { kind: 'Name', value: 'CreatePack' },
			variableDefinitions: [
				{
					kind: 'VariableDefinition',
					variable: { kind: 'Variable', name: { kind: 'Name', value: 'input' } },
					type: {
						kind: 'NonNullType',
						type: { kind: 'NamedType', name: { kind: 'Name', value: 'PackCreateRequest' } }
					}
				}
			],
			selectionSet: {
				kind: 'SelectionSet',
				selections: [
					{
						kind: 'Field',
						name: { kind: 'Name', value: 'packCreate' },
						arguments: [
							{
								kind: 'Argument',
								name: { kind: 'Name', value: 'input' },
								value: { kind: 'Variable', name: { kind: 'Name', value: 'input' } }
							}
						],
						selectionSet: {
							kind: 'SelectionSet',
							selections: [
								{
									kind: 'Field',
									name: { kind: 'Name', value: 'pack' },
									selectionSet: {
										kind: 'SelectionSet',
										selections: [
											{ kind: 'Field', name: { kind: 'Name', value: 'UID' } },
											{ kind: 'Field', name: { kind: 'Name', value: 'Size' } }
										]
									}
								}
							]
						}
					}
				]
			}
		}
	]
} as unknown as DocumentNode<CreatePackMutation, CreatePackMutationVariables>;
export const UpdatePackDocument = {
	kind: 'Document',
	definitions: [
		{
			kind: 'OperationDefinition',
			operation: 'mutation',
			name: { kind: 'Name', value: 'UpdatePack' },
			variableDefinitions: [
				{
					kind: 'VariableDefinition',
					variable: { kind: 'Variable', name: { kind: 'Name', value: 'input' } },
					type: {
						kind: 'NonNullType',
						type: { kind: 'NamedType', name: { kind: 'Name', value: 'PackUpdateRequest' } }
					}
				}
			],
			selectionSet: {
				kind: 'SelectionSet',
				selections: [
					{
						kind: 'Field',
						name: { kind: 'Name', value: 'packUpdate' },
						arguments: [
							{
								kind: 'Argument',
								name: { kind: 'Name', value: 'input' },
								value: { kind: 'Variable', name: { kind: 'Name', value: 'input' } }
							}
						],
						selectionSet: {
							kind: 'SelectionSet',
							selections: [
								{
									kind: 'Field',
									name: { kind: 'Name', value: 'pack' },
									selectionSet: {
										kind: 'SelectionSet',
										selections: [
											{ kind: 'Field', name: { kind: 'Name', value: 'UID' } },
											{ kind: 'Field', name: { kind: 'Name', value: 'Size' } }
										]
									}
								}
							]
						}
					}
				]
			}
		}
	]
} as unknown as DocumentNode<UpdatePackMutation, UpdatePackMutationVariables>;
export const DeletePackDocument = {
	kind: 'Document',
	definitions: [
		{
			kind: 'OperationDefinition',
			operation: 'mutation',
			name: { kind: 'Name', value: 'DeletePack' },
			variableDefinitions: [
				{
					kind: 'VariableDefinition',
					variable: { kind: 'Variable', name: { kind: 'Name', value: 'input' } },
					type: {
						kind: 'NonNullType',
						type: { kind: 'NamedType', name: { kind: 'Name', value: 'PackDeleteRequest' } }
					}
				}
			],
			selectionSet: {
				kind: 'SelectionSet',
				selections: [
					{
						kind: 'Field',
						name: { kind: 'Name', value: 'packDelete' },
						arguments: [
							{
								kind: 'Argument',
								name: { kind: 'Name', value: 'input' },
								value: { kind: 'Variable', name: { kind: 'Name', value: 'input' } }
							}
						],
						selectionSet: {
							kind: 'SelectionSet',
							selections: [
								{
									kind: 'Field',
									name: { kind: 'Name', value: 'pack' },
									selectionSet: {
										kind: 'SelectionSet',
										selections: [
											{ kind: 'Field', name: { kind: 'Name', value: 'UID' } },
											{ kind: 'Field', name: { kind: 'Name', value: 'Size' } }
										]
									}
								}
							]
						}
					}
				]
			}
		}
	]
} as unknown as DocumentNode<DeletePackMutation, DeletePackMutationVariables>;
export const CalculatePacksDocument = {
	kind: 'Document',
	definitions: [
		{
			kind: 'OperationDefinition',
			operation: 'query',
			name: { kind: 'Name', value: 'CalculatePacks' },
			variableDefinitions: [
				{
					kind: 'VariableDefinition',
					variable: { kind: 'Variable', name: { kind: 'Name', value: 'input' } },
					type: {
						kind: 'NonNullType',
						type: { kind: 'NamedType', name: { kind: 'Name', value: 'PackCalculateRequest' } }
					}
				}
			],
			selectionSet: {
				kind: 'SelectionSet',
				selections: [
					{
						kind: 'Field',
						name: { kind: 'Name', value: 'packCalculate' },
						arguments: [
							{
								kind: 'Argument',
								name: { kind: 'Name', value: 'input' },
								value: { kind: 'Variable', name: { kind: 'Name', value: 'input' } }
							}
						],
						selectionSet: {
							kind: 'SelectionSet',
							selections: [
								{
									kind: 'Field',
									name: { kind: 'Name', value: 'calculations' },
									selectionSet: {
										kind: 'SelectionSet',
										selections: [
											{ kind: 'Field', name: { kind: 'Name', value: 'PackUID' } },
											{ kind: 'Field', name: { kind: 'Name', value: 'PackSize' } },
											{ kind: 'Field', name: { kind: 'Name', value: 'Quantity' } }
										]
									}
								}
							]
						}
					}
				]
			}
		}
	]
} as unknown as DocumentNode<CalculatePacksQuery, CalculatePacksQueryVariables>;
