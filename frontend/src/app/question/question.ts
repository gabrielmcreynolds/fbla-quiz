export interface Question {
  question: string;
  choices?: Array<string>;
  correctChoice: string | boolean;
  selectedChoice?: string | boolean;
  // if the correctChoice is false this is the correct answer
  answer?: string;
}
