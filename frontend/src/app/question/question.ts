import { QuestionType } from './question-type.enum';

export interface Question {
  question: string;
  choices?: Array<string>;
  correctChoice: string | boolean;
  selectedChoice?: string | boolean;
  // if the correctChoice is false this is the correct answer
  answer?: string;
  type: QuestionType;

  isCorrect?: boolean;
}
