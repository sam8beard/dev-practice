import spacy 
import pprint
from spacy.matcher import Matcher

nlp = spacy.load("en_core_web_trf")

def lesson_1(): 
    
    doc = nlp("AI systems must be regulated to ensure safety.")

    for token in doc: 
        print(token.text, token.lemma_)
        print(token.pos_)
    
    for sent in doc.sents: 
        print(sent.text)

def lesson_2():
    doc = nlp("In 2023, the EU released a report arguing that OpenAI's models should be regulated.")
    for ent in doc.ents: 
        # print(ent.text, ent.label_)
        if ent.label_ == "ORG": 
            print(ent.text)

def lesson_3(): 
    # doc = nlp("The report argues that AI must be monitored. arguing.")
    # matcher = Matcher(nlp.vocab)

    # pattern = [{"LEMMA": "argue"}]
    # matcher.add("Argue", [pattern])

    # matches = matcher(doc)
    # my_matches = []
    # for match_id, start, end in matches: 
    #     span = doc[start:end]
    #     my_matches.append(span)

    # print(my_matches)
    doc = nlp("Scientists develop tools, and engineers build robots.")
    matcher = Matcher(nlp.vocab)

    pattern = [{"POS": "VERB"}, {"POS": "NOUN"}]

    matcher.add("ID", [pattern])
    matches = matcher(doc)
    for match_id, start, end in matches: 
        span = doc[start:end]
        print(span)

def lesson_4(): 
    doc = nlp("Some experts claim that AI will surpass human intelligence. Others argue it should be restricted.")
    matcher = Matcher(nlp.vocab)

    pattern1 = [{"ORTH": "claim"}]
    pattern2 = [{"ORTH": "argue"}]

    matcher.add("Claim", [pattern1])
    matcher.add("Argue", [pattern2])

    matches = matcher(doc)
    sentences = []
    for match_id, start, end in matches: 
        sentences.append(doc[start].sent)


    print(sentences)

def lesson_5(): 
    doc = nlp(
        "Critics argue that regulation is necessary. " \
        "Proponents suggest it should be limited. " \
        "Researchers show evidence that oversight improves safety. " \
        "Don't go back to Rockville. " \
        "What's the frequency Kenneth?"
        )
    matcher = Matcher(nlp.vocab)
    # print(doc)
    pattern1 = [{"POS": "VERB", "LEMMA": {"IN": ["argue", "suggest"]}}]

    pattern2 = [{"LEMMA": "researcher"}, {"LEMMA": "show"}]

    matcher.add("ArgueSuggest", [pattern1])
    matcher.add("ResearchShows", [pattern2])

    matches = matcher(doc)

    # Must extract sentences that are claims 

    claim_sentences = []
    for match_id, start, end in matches:
        string_id = nlp.vocab.strings[match_id]
        # print(match_id, string_id, start, end)
        claim_sentences.append(doc[start].sent)

    print(claim_sentences)

def lesson_6(): 
    doc = nlp("In 2022, the UN stated that AI must be controlled. " \
    "Some experts argue that ethical guidelines are necessary. " \
    "Others suggest innovation should not be restricted."
    )

    matcher = Matcher(nlp.vocab)
    claim_verbs = ["state", "argue", "suggest"]
    pattern = [{"LEMMA": {"IN": claim_verbs}}]
    matcher.add("Claims", [pattern])

    matches = matcher(doc)
    claim_sentences = [] 
    
    claims_by_orgs = {}
    print(doc)
    for match_id, start, end in matches:
        # print(doc[start].sent)
        print(doc[start].sent.ents)
        orgs = [ent for ent in doc[start].sent.ents if ent.label_ == "ORG"]
        if orgs: 
           claims_by_orgs[tuple(orgs)] = doc[start].sent
        claims_by_orgs["No organizational entity"] = doc[start].sent 
        
    pprint.pprint(claims_by_orgs)
# lesson_1()
# lesson_2()
# lesson_3()
# lesson_4()
# lesson_5()
lesson_6()