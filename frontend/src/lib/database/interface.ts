import {Source} from '../models/database/source';
// import {SourceSummary} from '../../app/models/fasten/source-summary';

export interface IDatabaseDocument {
  populateId(): void
}

export interface IDatabasePaginatedResponse {
  offset: number,
  total_rows: number,
  rows: any[]
}
export interface IDatabaseRepository {
  GetDB(): any
  Close(): void

  // CreateUser(*models.User) error
  // GetUserByEmail(context.Context, string) (*models.User, error)
  // GetCurrentUser(context.Context) *models.User

  // GetSummary(ctx context.Context) (*models.Summary, error)

  // UpsertResource(context.Context, *models.ResourceFhir) error
  // GetResourceBySourceType(context.Context, string, string) (*models.ResourceFhir, error)
  // GetResourceBySourceId(context.Context, string, string) (*models.ResourceFhir, error)
  // ListResources(context.Context, models.ListResourceQueryOptions) ([]models.ResourceFhir, error)
  // GetPatientForSources(ctx context.Context) ([]models.ResourceFhir, error)

  CreateSource(source: Source): Promise<string>
  GetSource(source_id: string): Promise<Source>
  DeleteSource(source_id: string): Promise<boolean>
  // GetSourceSummary(source_id: string): Promise<SourceSummary>
  GetSources(): Promise<IDatabasePaginatedResponse>
}
