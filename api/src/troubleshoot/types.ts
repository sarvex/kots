
export interface SupportBundle {
  id: string;
  slug: string;
  watchId: string;
  name: string;
  size: number;
  notes: string;
  status: string;
  uri: string;
  resolution: string;
  treeIndex: string;
  viewed: boolean;
  createdAt: string;
  uploadedAt: string;
  isArchived: boolean;

  signedUri: string;
  analysis: SupportBundleAnalysis;
  watchSlug: string;
  watchName: string;
};

export interface SupportBundleAnalysis {
  id: string;
  error: string;
  maxSeverity: string;
  insights: SupportBundleInsight[];
  createdAt: string;
};

export interface SupportBundleInsight {
  key: string;
  severity: string;
  primary: string;
  detail: string;
  icon: string;
  icon_key: string;
  desiredPosition: number;
  labels: Label[];
}

export interface Label {
  key: string;
  value: string;
}
