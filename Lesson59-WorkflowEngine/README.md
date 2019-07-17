
state=draft plan; event=next;
        => state=matchusers; code:{look at scope (*config) and update MatchingUsers[]}; event=next
            =>> state=assignuser; code{use assignuser method}; event=next
                => state=review; 
                        if event="approve" {
                            => state=createplan; code{flow that creates a plan}; event=next
                        }
                        else if event="deny" {
                            => state=_end
                        }


cnDevelopService:

1) Request:started 
    event="next"
2)    => Request:wip
        event="next"
3)         Request:done
            event="next"
4)              Rule:started
                    *fn => getting list of users based on scope and assigning that to matched users
                    event="next"
5)                  Rule:wip
                     t=n; upon file upload; event="next"
6)                      Rule:done 

7)                          Code:started
                                *fn => getting list of users based on scope and assigning that to matched users



=========


cnwebserver -> cndevelopservice  (myworkspace and rule request dashboard)
    Rule req dashboard:
    1) give me all the flowstates where request.user = me

    myworkspace:
    1) give me all the flow states where assigned user = me
        check-in:
            => "next"
