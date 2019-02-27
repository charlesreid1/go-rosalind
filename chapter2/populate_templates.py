import jinja2
import os

def main():

    # Jinja env
    env = jinja2.Environment(loader=jinja2.FileSystemLoader('.'))
    
    problems = [
        {
            'chapter': '2',
            'problem': 'a',
            'title': 'Implement Motif Enumeration',
            'description': 'Given a collection of strings of DNA, find all motifs (kmers of length k and Hamming distance d from all DNA strings).',
            'url': 'http://rosalind.info/problems/ba2a/'
        },
        {
            'chapter': '2',
            'problem': 'b',
            'title': 'Find a Median String',
            'description': 'Given a kmer length k and a set of strings of DNA, find the kmer(s) that minimize the L1 norm of the distance from it to all other DNA strings.',
            'url': 'http://rosalind.info/problems/ba2b/'
        },
        {
            'chapter': '2',
            'problem': 'c',
            'title': 'Find a Profile-most Probable k-mer in a String',
            'description': 'Given a profile matrix, find the most probable k-mer to generate the given DNA string.',
            'url': 'http://rosalind.info/problems/ba2c/'
        },
        {
            'chapter': '2',
            'problem': 'd',
            'title': 'Implement GreedyMotifSearch',
            'description': 'Find a collection of motif strings using a greedy motif search. Return first-occurring profile-most probable kmer.',
            'url': 'http://rosalind.info/problems/ba2d/'
        },
        {
            'chapter': '2',
            'problem': 'e',
            'title': 'Implement GreedyMotifSearch with Pseudocounts',
            'description': 'Re-implement problem BA2d (greedy motif search) using pseudocounts, which avoid setting probabilities to an absolute value of zero.',
            'url': 'http://rosalind.info/problems/ba2e/'
        },
    ]
    
    print("Writing problem boilerplate code")
    
    t = 'template.go.j2'
    for problem in problems:
        contents = env.get_template(t).render(**problem)
        fname = 'ba'+problem['chapter']+problem['problem']+'.go'
        if not os.path.exists(fname):
            print("Writing to file %s..."%(fname))
            with open(fname,'w') as f:
                f.write(contents)
        else:
            print("File %s already exists, skipping..."%(fname))
    
    print("Done")

if __name__=="__main__":
    main()
