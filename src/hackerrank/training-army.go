// In the magical kingdom of Kasukabe, people strive to possess skillsets. Higher the number of skillset present among the people, the more content people will be.

// There are N types of skill set present and initially there exists Ci people possessing ith skill set, where i in [1,N].

// There are T wizards in the kingdom and they have the ability to transform the skill set of a person into another skill set. Each of the these wizards has two lists of skill sets associated with them, A and B. He can only transform the skill set of person whose initial skill set belongs to the list A to one of the final skill set which belongs to the list B. That is, if A = [2,3,6] and B=[1,2] then following transformation can be done by that trainer.
// 2 -> 1
// 2 -> 2
// 3 -> 1
// 3 -> 2
// 6 -> 1
// 6 -> 2

// Once a transformation is done, both skill is removed from the respective lists. In the above example, if he perform 3->1 transformation on a person, list A will be updated to [2,6] and list B will be [2]. This updated list will be used for further transformations.

// Few points to note are:

// One person can possess only one skill set.
// A wizard can perform zero or more transformation as long as they satisfies the above criteria.
// A person can go through multiple transformation of skill set.
// Same class transformation is also possible. That is a person' skill set can be transformed into his current skill set. Eg.  in the above example.
// Your goal is to design a series of transformation which results into maximum number of skill set with non-zero number of people knowing it.

// Input Format

// The first line contains two numbers,N T, where N represent the number of skill set and T represent the number of wizards.
// Next line contains N space separated integers,C1 C2.... Cn, where Ci represents the number of people with ith skill. Then follows 2 x T lines, where each pair of line represent the configuration of each wizard.
// First line of the pair will start with the length of list A and followed by list B in the same line. Similarly second line of the pair starts with the length of list B and then the list B.

// Constraints
//  1 <= N <= 200
//  0 <= T <= 30
//  0 <= Ci <= 10
//  0 <= |A| <= 50
//  1 <= Ai <= N
//  Ai # Aj, 1 <= i <= j <= |A|
//  0 <= |B| <= 50
//  1 <= Bi <= N
//  Bi # Bj, i <= i < = j <= |B|

// Output Format

// The output must consist of one number, the maximum number of distinct skill set that can the people of country learn, after making optimal transformation steps.

// Sample Input

// 3 3
// 3 0 0
// 1 1
// 2 2 3
// 1 2
// 1 3
// 1 1
// 1 2 
// Sample Output

// 2  
// Explanation

// There are 3 types of skill sets present along with 3 wizards. Initially, all three people know the 1st skill set but no one knows the 2nd and 3rd skill sets.

// The 1st wizard's initial lists are: A=[1] and B=[2,3]. Suppose, he performs 1->2 transformation one any one of person with the 1st skill set, then it's list A will be updated to an empty list [] and list B will be [3].
// Now, we have two people knowing the 1st skill set and one person knowing the 2nd skill set.

// The 3rd wizard's initial lists are: A=[1] and B=[2]. He will use the transformation 1->2 one of the person with the 1st skill set, then it's lists will also be updated to an empty lists A:[] and B:[] .

// Now, we have 1 person with 1st skillset and and 2 people knowing the 2nd skillset.

// The 2nd wizard's initial lists are:A=[2]  and B=[3]. He will transform one of the person with 2nd skillset to 3rd one using the transformation 2->3. It's lists will also be updated to an empty lists A:[] and B:[].
// At this point, no further transformations are possible and we have achieved our maximum possible answer. Thus, each of the skill set, is known by 1 person.. This means there are three skill sets available in the kingdom.

package hackerrank

func TrainingArmy(){

}

func trainingarmy_solution(){

}